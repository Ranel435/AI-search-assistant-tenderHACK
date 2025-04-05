import React, { useState } from "react";
import ChatWindow from "./components/ChatWindow/ChatWindow";
import ChatInput from "./components/ChatInput/ChatInput";
import "./App.css";

function App() {
  const [messages, setMessages] = useState([]);
  const [isFirstMessageSent, setIsFirstMessageSent] = useState(false);

  const sendMessage = (text) => {
    if (!text.trim()) return;

    setIsFirstMessageSent(true);

    const userMessage = { text, sender: "user" };
    const botResponse = { text: "Ваш аккаунт qquerell заблокирован Причина блокировки: подозрительные действия с вашей стороны. Время от времени на нашем сайте появляются мошенники, поэтому сотрудникам приходится быть крайне бдительными и работать на опережение. Мы хотим разобраться в ситуации, поэтому, пожалуйста, напишите нам. Обычно отвечаем за 2 дня, но постараемся быстрее. Просим вас не писать каких-либо отзывов и не начинать публичное обсуждение до окончания разбирательства. Нам важно поддерживать порядок на площадке, но, к сожалению, добиться этого без блокировок не представляется возможным", sender: "bot" };

    setMessages((prev) => [...prev, userMessage, botResponse]);
  };

  return (
    <div className="app">
      <main className="app__container">
        <ChatWindow messages={messages} />
      </main>
      <ChatInput send={sendMessage} isFirstMessageSent={isFirstMessageSent} />
    </div>
  );
}
export default App;


// import React, { useState, useEffect } from "react";
// import ChatWindow from "./components/ChatWindow/ChatWindow";
// import ChatInput from "./components/ChatInput/ChatInput";
// import ChatList from "./components/ChatList/ChatList";  // Добавляем список чатов
// import "./App.css";

// function App() {
//   const [chats, setChats] = useState({});  // Все чаты
//   const [activeChat, setActiveChat] = useState(null);  // Текущий чат

//   const sendMessage = (text) => {
//     if (!text.trim() || !activeChat) return;

//     const userMessage = { text, sender: "user" };
//     const botResponse = { text: "Ответ бота", sender: "bot" };

//     setChats((prevChats) => ({
//       ...prevChats,
//       [activeChat]: [...(prevChats[activeChat] || []), userMessage, botResponse]
//     }));
//   };

//   const createNewChat = () => {
//     const newChatId = `chat_${Date.now()}`;
//     setChats((prevChats) => ({ ...prevChats, [newChatId]: [] }));
//     setActiveChat(newChatId);
//   };

//   useEffect(() => {
//     const savedChats = localStorage.getItem("chats");
//     if (savedChats) {
//       setChats(JSON.parse(savedChats));
//     }
//   }, []);
  
//   useEffect(() => {
//     localStorage.setItem("chats", JSON.stringify(chats));
//   }, [chats]);
  
//   return (
//     <div className="app">
//       <ChatList chats={chats} activeChat={activeChat} setActiveChat={setActiveChat} createNewChat={createNewChat} />
//       <main className="app__container">
//         {activeChat ? (
//           <>
//             <ChatWindow messages={chats[activeChat] || []} />
//             <ChatInput send={sendMessage} />
//           </>
//         ) : (
//           <div className="empty-chat">Выберите или создайте новый чат</div>
//         )}
//       </main>
//     </div>
//   );
// }

// export default App;
