import React from "react"
import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from './pages/Home/Home';
import Sales from "./pages/Sales/sales";

function App() {
  return (
    // <div className="App">
    //   <header className="App-header">
    //     <Home></Home>
        
    //   </header>
    // </div>

   <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home/>}/>
      <Route path="/sales/add" element={<Sales/>}/>
    </Routes>
   </BrowserRouter>
  );
}

// export default function App() {
//   return <h1>Hello World</h1>;
// }

export default App;
