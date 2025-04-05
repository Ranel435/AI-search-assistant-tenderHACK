from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
import os

# Путь к локальной модели (используем абсолютный путь или относительный с ./*)
model_path = "/Users/dremotha/Projects/GolangProjects/tenderhack/llm_services/llm-model"
if not os.path.exists(model_path):
    raise ValueError(f"Путь {model_path} не существует. Проверьте расположение модели.")

# Загрузка токенизатора и модели из локальной директории
tokenizer = AutoTokenizer.from_pretrained(model_path, local_files_only=True)
model = AutoModelForCausalLM.from_pretrained(model_path, local_files_only=True)

# Подготовка данных для инференса
question = "Какие этапы включает процесс работы с контрактами?"

# Токенизация вопроса
inputs = tokenizer(question, return_tensors="pt")

# Выполнение инференсао
with torch.no_grad():
    outputs = model.generate(
        input_ids=inputs.input_ids,
        attention_mask=inputs.attention_mask,
        max_new_tokens=50,  # Максимальное количество новых токенов
        do_sample=True,  # Использовать сэмплирование для генерации
        top_k=50,  # Использовать top-k sampling
        top_p=0.95,  # Использовать nucleus sampling
    )

# Декодирование и вывод ответа
response = tokenizer.decode(outputs[0], skip_special_tokens=True)
print("Ответ:", response)