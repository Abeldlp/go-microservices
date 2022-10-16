import { useParams } from "react-router-dom";

const ShopDetails: React.FC = () => {
  let { id } = useParams();
  return (
    <>
      <h1>Shop main view</h1>
      <p>Hello {id}</p>
    </>
  );
};

export default ShopDetails;
