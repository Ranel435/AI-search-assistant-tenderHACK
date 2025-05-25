import React from "react";
import styles from "../ChatWindow/ChatWindow.module.css";

class ChatMessage extends React.Component {
  render() {
    
      useEffect(() => {
        if (msg.sender === "bot") {
          setIsLoading(true);
          // Имитация задержки ответа 2 секунды
            setIsLoading(false);
          }, 2000);
        return  () => clearTimeout(timer);
        }
      }, [msg.sender]);
    
        return  (
        <div
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
            {isLoading ? (
              <div className={styles.searching}>
                <span>Ищем...</span>
                <div className={styles.loadingDots}>
                  <div className={styles.dot}></div>
                  <div className={styles.dot}></div>
                  <div className={styles.dot}></div>
                </div>
              </div>
            ) : (
              <>
                <span className="f24">{msg.text}</span>
                {msg.sender === "bot" && (
                  <div className={styles.message__feedback}>
                    <img className={styles.message__feedback_img} src="/dislike.svg" alt="Dislike" />
                    <img className={styles.message__feedback_img}  src="/like.svg" alt="Like" />
                  </div>
                )}
              </>
            )}
          </div>
        </div>
      );
    }
    
    export default ChatMessage;
    
  }
}

export default ChatMessage;
