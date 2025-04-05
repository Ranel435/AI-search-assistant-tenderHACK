import React, { useEffect, useRef } from "react";
import styles from "./ChatWindow.module.css";

function ChatWindow({ messages }) {
  const chatRef = useRef(null);

  useEffect(() => {
    if (chatRef.current) {
      chatRef.current.scrollTop = chatRef.current.scrollHeight;
    }
  }, [messages]);
  return (
    <div className={styles.window} ref={chatRef}>
      {messages && messages.length > 0 ? (
        messages.map((msg, index) => (
          <div
            key={index}
            className={`${styles.message} ${
              msg.sender === "user" ? styles.userMessage : styles.botMessage
            }`}
          >
            <img
              src="/chatIcon.svg"
              alt=""
              className={`${styles.message__img} ${
                msg.sender === "bot" ? styles.active : styles.nonactive
              }`}
            />
            <div className={`${styles.messageContent} f16`}>
              <span>{msg.text}</span>

              {msg.sender === "bot" && (
                <div className={styles.message__feedback}>
                  <img src="/dislike.svg" alt="Dislike" />
                  <img src="/like.svg" alt="Like" />
                </div>
              )}
            </div>
          </div>
        ))
      ) : (
        <div className={styles.nomessages}>
          <h1 className={`f30 ${styles.noMessages__title}`}>
            <img src="/chatIcon.svg" /> Чат-бот портала поставщиков
          </h1>
          <h2 className={`f16 ${styles.noMessages__subtitle}`}>
            Как я могу помочь?
          </h2>
        </div>
      )}
    </div>
  );
}

export default ChatWindow;
