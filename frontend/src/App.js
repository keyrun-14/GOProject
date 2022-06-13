import {BrowserRouter, Route,Routes} from "react-router-dom";
import './App.css';
import AddProduct from "./components/AddProduct/AddProduct";

import Register from "./components/register/Register";
import Login from "./components/login/Login"
import Home from "./components/home/Home"
import Products from "./components/Products/Products"
import EditProduct from "./components/editproduct/EditProduct";

function App() {
  return (
    <div className="App">
          
      <BrowserRouter>
        <Routes>
        
          <Route path="/" element={<Home />} />    
          <Route exact path="/register" element={<Register />} />
          <Route exact path="/login" element={<Login />} />
          <Route exact path="/products" element={<Products />} />
        
          <Route path="/edit" element={<EditProduct />} />
       
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
