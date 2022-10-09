import "./App.css";
import MainContent from "./Components/Content/MainContent";
import MainLayout from "./Components/Layouts/MainLayout";
import SideNavitation from "./Components/Navigation/SideNavitation";

function App() {
  return (
    <div className="App">
      <MainLayout left={<SideNavitation />} center={<MainContent />} />
    </div>
  );
}

export default App;
