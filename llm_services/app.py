from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
from typing import Optional

app = FastAPI()

# Модель будет загружена при старте сервера
print("Загрузка модели...")
model_path = "./llm-model" 
tokenizer = AutoTokenizer.from_pretrained(model_path)

# Настройки для экономии памяти
model = AutoModelForCausalLM.from_pretrained(
    model_path,
    torch_dtype=torch.float16,
    low_cpu_mem_usage=True
)
print("Модель загружена!")

class GenerationRequest(BaseModel):
    prompt: str
    context: Optional[str] = ""
    max_tokens: Optional[int] = 512

@app.post("/generate")
async def generate(request: GenerationRequest):
    try:
        # Формирование промпта
        if request.context:
            full_prompt = f"Контекст: {request.context}\n\nВопрос: {request.prompt}\n\nОтвет:"
        else:
            full_prompt = f"Вопрос: {request.prompt}\n\nОтвет:"
        
        # Токенизация
        inputs = tokenizer(full_prompt, return_tensors="pt")
        
        # Генерация ответа
        with torch.no_grad():
            outputs = model.generate(
                input_ids=inputs.input_ids,
                attention_mask=inputs.attention_mask,
                max_new_tokens=request.max_tokens,
                do_sample=True,
                temperature=0.7,
                top_k=50,
                top_p=0.95
            )
        
        # Декодирование результата
        full_output = tokenizer.decode(outputs[0], skip_special_tokens=True)
        
        # Выделяем только ответ (без повторения вопроса)
        response = full_output[len(full_prompt):].strip()
        
        return {"generated_text": response}
    
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# Проверка работоспособности
@app.get("/health")
async def health_check():
    return {"status": "ok"}

# Запуск на всех интерфейсах
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
