import React, { useState } from "react";
import styles from './ChatInput.module.css';

function ChatInput({ send }) {
  const [input, setInput] = useState("");

  const handleSend = () => {
    if (!input.trim()) return;
    send(input);
    setInput("");
  };

  return (
    <div className={styles.chatInput}>
      <input
        className={` f16 ${styles.chatInput__input}`}
        type="text"
        placeholder="Задайте вопрос"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        onKeyPress={(e) => e.key === "Enter" && handleSend()}
      />
      <button className={` f16 ${styles.chatInput__button}`} onClick={handleSend}>
        Отправить
      </button>
    </div>
  );
}

export default ChatInput;

