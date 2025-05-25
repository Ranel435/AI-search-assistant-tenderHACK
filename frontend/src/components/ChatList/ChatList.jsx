import React from "react";
import styles from "../ChatWindow/ChatWindow.module.css";

class ChatList extends React.Component {
  render() {
    
        return  (
        <div className={styles.chatList}>
          <div className={styles.chatList__container}>
            <button className={` f16 ${styles.newChatBtn}`} onClick={createNewChat}>
              <img src="/newChat.svg" alt="" /> Новый чат
            </button>
            <ul>
              {Object.keys(chats).map((chatId) => (
                <li
                  key={chatId}
                  className={`${styles.chatName} ${chatId === activeChat ? styles.active : ""}`}
                  onClick={() => setActiveChat(chatId)}
                >
                  {chatId.split("_")[1]}
                </li>
              ))}
            </ul>
          </div>
        </div>
      );
    }
    
    export default ChatList;
  }
}

export default ChatList;
