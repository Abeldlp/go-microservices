import "./MainContent.css";
import { Routes, Route } from "react-router-dom";
import Home from "../../views/Home";
import ShopDetails from "../../views/ShopDetails";

const MainContent: React.FC = () => {
  return (
    <div className="main-content-container">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/:id" element={<ShopDetails />} />
      </Routes>
    </div>
  );
};

export default MainContent;
