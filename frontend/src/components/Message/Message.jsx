import React from "react";
import styles from "../ChatWindow/ChatWindow.module.css";

class Message extends React.Component {
  render() {
        return  (
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
  }
}

export default Message;
