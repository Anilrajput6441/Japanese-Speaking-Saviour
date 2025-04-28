import { useEffect, useState } from "react";

function App() {
  const [socket, setSocket] = useState(null);
  const [input, setInput] = useState("");
  const [translated, setTranslated] = useState("");
  const [audioUrl, setAudioUrl] = useState("");
  const [debouncedInput, setDebouncedInput] = useState(input);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setTranslated(data.translated);
      setAudioUrl(data.audioUrl);
    };
    setSocket(ws);

    return () => ws && ws.close();
  }, []);

  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedInput(input);
    }, 500); // 500ms debounce time

    return () => clearTimeout(timer); // Cleanup on component unmount or input change
  }, [input]);

  useEffect(() => {
    if (socket && debouncedInput) {
      socket.send(JSON.stringify({ text: debouncedInput }));
    }
  }, [debouncedInput, socket]);

  const handleChange = (e) => {
    setInput(e.target.value);
  };

  return (
    <div className="relative App flex flex-col justify-center items-center w-[100vw] h-[100vh] ">
      <img
        className="absolute h-[100%] w-[100%]"
        src="https://img.freepik.com/free-vector/hand-drawn-japanese-illustration-cherry-tree-petals_23-2149601832.jpg?t=st=1745877505~exp=1745881105~hmac=0eff20ae38afa4bdc41ac3804a7b1911d61cc285bcea4a9ed9c8223e1c462495&w=2000"
        alt=""
      />
      <div className="absolute text-center flex justify-center items-center flex-col ">
        <h1 className="font-bold text-6xl p-5">Live Japanese Translator</h1>
        <textarea
          value={input}
          onChange={handleChange}
          placeholder="Type or Write..."
          className="border-2 w-[90%] p-2 h-15 outline-none"
        />
        <h2 className="font-bold bg-white w-[90%] mt-5">Translation:</h2>
        <p className="bg-white w-[90%]">{translated}</p>

        {audioUrl && (
          <audio key={audioUrl} controls className="mt-5">
            <source src={audioUrl} type="audio/mpeg" />
            Your browser does not support the audio element.
          </audio>
        )}
      </div>
    </div>
  );
}

export default App;
