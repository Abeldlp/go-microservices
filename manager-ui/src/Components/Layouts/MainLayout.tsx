import { ReactNode } from "react";
import "./MainLayout.css";

type Node = ReactNode | ReactNode[];

interface Props {
  left: Node;
  center: Node;
}

const MainLayout: React.FC<Props> = (props: Props) => {
  return (
    <div className="main-layout">
      <div>{props.left}</div>
      <div>{props.center}</div>
    </div>
  );
};

export default MainLayout;
