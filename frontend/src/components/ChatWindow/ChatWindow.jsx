import React from "react";
import styles from "../ChatWindow/ChatWindow.module.css";

class ChatWindow extends React.Component {
  render() {
      // скролл к текущему сообщению
        if (chatRef.current) {
          chatRef.current.scrollTop = chatRef.current.scrollHeight;
        }
      }
      useEffect(() => {
        // if (chatRef.current) {
        //   chatRef.current.scrollTop = chatRef.current.scrollHeight;
        // }
        scroll();
      }, [messages]);
      // название чата
    
        return  (
        <div className={styles.chatContainer}>
          <div className={styles.chatHeader}>{chatTitle}</div>
    
          <div className={styles.window} ref={chatRef}>
            {messages && messages.length > 0 ? (
              messages.map((msg, index) => (
                <ChatMessage key={index} msg={msg} scroll={scroll}/>
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
        </div>
      );
    }
    
    export default ChatWindow;
    
  }
}

export default ChatWindow;
