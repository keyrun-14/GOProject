import React from "react";
import "../home/home.css"
import NavBar from "../Nav/NavBar"
const Home = () => {
  return (
    <>
     <NavBar />
      <div className="main">
        <div className="heading">
          <p className="welcome">
            <span style={{ color: "#66CC66" }}>W</span>
            <span style={{ color: "#FF9966" }}>E</span>
            <span style={{ color: "#5c3488" }}>L</span>
            <span style={{ color: "#FF9966" }}>C</span>
            <span style={{ color: "#5c3488" }}>O</span>
            <span style={{ color: "#66CC66" }}>M</span>
            <span style={{ color: "#5c3488" }}>E</span>
          </p>
          {/* <h1 className="app-name">
            <span style={{ color: "#66CC66" }}>Photo</span>
            <span style={{ color: "#FF9966" }}>Sharing</span>
            <span style={{ color: "#5c3488" }}>App</span>
          </h1> */}
        </div>
        {/* <div className="blocks">
          <div className="box share">Share</div>
          <div className="box your">your</div>
          <div className="box story"> story</div>
        </div> */}
      </div>
    </>
  );
};

export default Home;
