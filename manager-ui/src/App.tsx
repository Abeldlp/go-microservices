import "./App.css";
import MainContent from "./Components/Content/MainContent";
import MainLayout from "./Components/Layouts/MainLayout";
import SideNavitation from "./Components/Navigation/SideNavitation";
import { BrowserRouter as Router } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Router>
        <MainLayout left={<SideNavitation />} center={<MainContent />} />
      </Router>
    </div>
  );
}

export default App;
