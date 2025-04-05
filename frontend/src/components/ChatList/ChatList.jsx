import React from "react";
import styles from "./ChatList.module.css";

function ChatList({ chats, activeChat, setActiveChat, createNewChat }) {
  return (
    <div className={styles.chatList}>
      <button className={styles.newChatBtn} onClick={createNewChat}>➕ Новый чат</button>
      <ul>
        {Object.keys(chats).map((chatId) => (
          <li 
            key={chatId} 
            className={chatId === activeChat ? styles.active : ""}
            onClick={() => setActiveChat(chatId)}
          >
            Чат {chatId.split("_")[1]}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ChatList;
