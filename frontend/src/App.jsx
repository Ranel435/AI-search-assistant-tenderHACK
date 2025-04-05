import React, { useState, useEffect } from "react";
import ChatWindow from "./components/ChatWindow/ChatWindow";
import ChatInput from "./components/ChatInput/ChatInput";
import ChatList from "./components/ChatList/ChatList";  
import "./App.css";

function App() {
  const [chats, setChats] = useState(() => {
    const savedChats = localStorage.getItem("chats");
    return savedChats ? JSON.parse(savedChats) : {}; // Загружаем чаты при первой загрузке
  });

  const [activeChat, setActiveChat] = useState(() => {
    return localStorage.getItem("activeChat") || null; // Загружаем последний активный чат
  });

  useEffect(() => {
    localStorage.setItem("chats", JSON.stringify(chats));
  }, [chats]);

  useEffect(() => {
    if (activeChat) {
      localStorage.setItem("activeChat", activeChat);
    }
  }, [activeChat]);

  const sendMessage = (text) => {
    if (!text.trim() || !activeChat) return;

    const userMessage = { text, sender: "user" };
    const botResponse = { text: "gnw[ojjgn]oqwn", sender: "bot"};
    // const botResponse = { text: "We’ve trained a model called ChatGPT which interacts in a conversational way. The dialogue format makes it possible for ChatGPT to answer followup questions, admit its mistakes, challenge incorrect premises, and reject inappropriate requests.ChatGPT is a sibling model to InstructGPT⁠, which is trained to follow an instruction in a prompt and provide a detailed response.We are excited to introduce ChatGPT to get users’ feedback and learn about its strengths and weaknesses. During the research preview, usage of ChatGPT is free. Try it now at chatgpt.com⁠(opens in a new window).", sender: "bot" };

    setChats((prevChats) => ({
      ...prevChats,
      [activeChat]: [...(prevChats[activeChat] || []), userMessage, botResponse],
    }));
  };

  const createNewChat = () => {
    const now = new Date();
    const day = String(now.getDate()).padStart(2, '0');
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const year = now.getFullYear();
    const time = now.getTime();
  
    const formattedDate = `${day}.${month}.${year}`;
    const newChatId = `chat_${formattedDate}_${time}`;
  
    setChats((prevChats) => ({ ...prevChats, [newChatId]: [] }));
    setActiveChat(newChatId);
  };
  

  return (
    <div className="app">
      <ChatList
        chats={chats}
        activeChat={activeChat}
        setActiveChat={setActiveChat}
        createNewChat={createNewChat}
      />
      <main className="app__container">
        {activeChat ? (
          <>
            <ChatWindow messages={chats[activeChat] || []} />
            <ChatInput send={sendMessage} />
          </>
        ) : (
          <div className="empty-chat">Выберите или создайте новый чат</div>
        )}
      </main>
    </div>
  );
}

export default App;
