import React from "react";
import styles from "./Message.module.css";

const Message = ({ sender, text }) => {
  return (
    <div className={`${styles.messageContainer} ${sender === "user" ? styles.userContainer : styles.botContainer}`}>
      {sender === "bot" && (
        <img 
          src="/chatIcon.svg" 
          alt="Bot" 
          className={styles.messageAvatar}
        />
      )}
      
      <div className={`${styles.messageContent} ${sender === "user" ? styles.userMessage : styles.botMessage}`}>
        {text}
      </div>
    </div>
  );
};

export default Message;