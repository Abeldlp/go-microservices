import "./MainContent.css";
import { Routes, Route } from "react-router-dom";
import Home from "../../views/Home";
import ShopDetails from "../../views/ShopDetails";
import Settings from "../../views/Settings";

const MainContent: React.FC = () => {
  return (
    <div className="main-content-container">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/shops/:id" element={<ShopDetails />} />
        <Route path="/settings/:id" element={<Settings />} />
      </Routes>
    </div>
  );
};

export default MainContent;
