import { useParams } from "react-router-dom";

const Settings: React.FC = () => {
  let { id } = useParams();

  return (
    <>
      <h1>Settings main view</h1>
      <p>Settings for {id}</p>
    </>
  );
};

export default Settings;
