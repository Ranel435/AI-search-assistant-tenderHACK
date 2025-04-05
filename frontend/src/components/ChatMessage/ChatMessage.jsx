// import React from "react";
// import styles from "../ChatWindow/ChatWindow.module.css";

// function ChatMessage({ msg }) {
//   return (
//     <div
//       className={`${styles.message} ${
//         msg.sender === "user" ? styles.userMessage : styles.botMessage
//       }`}
//     >
//       <img
//         src="/chatIcon.svg"
//         alt=""
//         className={`${styles.message__img} ${
//           msg.sender === "bot" ? styles.active : styles.nonactive
//         }`}
//       />
//       <div className={`${styles.messageContent} f16`}>
//         <span>{msg.text}</span>

//         {msg.sender === "bot" && (
//           <div className={styles.message__feedback}>
//             <img src="/dislike.svg" alt="Dislike" />
//             <img src="/like.svg" alt="Like" />
//           </div>
//         )}
//       </div>
//     </div>
//   );
// }

// export default ChatMessage;


import React, { useState, useEffect } from "react";
import styles from "../ChatWindow/ChatWindow.module.css";

function ChatMessage({ msg , scroll}) {
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (msg.sender === "bot") {
      setIsLoading(true);
      // Имитация задержки ответа 2 секунды
      const timer = setTimeout(() => {
        setIsLoading(false);
      }, 2000);
      return () => clearTimeout(timer);
    }
  }, [msg.sender]);

  return (
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

