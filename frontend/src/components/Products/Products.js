import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import AddProduct from "../AddProduct/AddProduct";
import EditProduct from "../editproduct/EditProduct";
import "./Products.css"

const Products = () => {
  const [data, setData] = useState([]);
  const [seen, setSeen] = useState(false)
  const [productId, setProductId] = useState(0)

  useEffect(() => {
    fetch("http://localhost:8080/products", {
      method: "GET",
      mode: "cors",
      headers: {
        "content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
      },
    })
      .then((res) => res.json())
      .then((datas) => setData(datas))

  }, []);
  // console.log(data);
  // console.log(data.length);
  function deletepost(id) {
    fetch("http://localhost:8080/product/" + id, {
      method: "DELETE",
      mode: "cors",
      headers: {
        "content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
      },
    })
      .then((res) => console.log(res))
      .catch((e) => console.log(e));
      window.location.href="http://localhost:3000/products"
  }
  function changeSeen(id){
    setProductId(id)
    setSeen(true)
    console.log(seen)
  }
  return (
    <>
  <AddProduct/>
    {  data==null ?
     (<table><tr>
  <th>product</th>
  <th>model</th>
  <th>specification</th>
  <th>price</th>
</tr></table>) : (
<table>
  <thead>
  <tr>
  <th>product</th>
  <th>model</th>
  <th>specification</th>
  <th>price</th>
</tr>

  </thead>

{data.map((item) => (
  <tbody>
    <tr>
    <td> {item.name}</td>
    <td>{item.model}</td>
    <td>{item.specification}</td>
    <td>{item.price}</td>
  <td> <i class="fa-solid fa-pen" onClick={()=>changeSeen(item.id)}></i></td> 
  
   
    <td><i onClick={()=>deletepost(item.id)} class="fa-solid fa-trash-can"></i></td>
  </tr>

  </tbody>
  
))}
</table>)
      }
    <EditProduct seenseen={seen} productId={productId} />
      
    
    </>
  );
};

export default Products;
