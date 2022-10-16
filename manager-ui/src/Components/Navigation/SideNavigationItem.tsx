import { ListItem, ListItemButton, ListItemText } from "@mui/material";
import { Link } from "react-router-dom";

interface Props {
  text: string;
  to: string;
}

const SideNavigationItem: React.FC<Props> = (props: Props) => {
  return (
    <>
      <Link to={props.to}>
        <ListItem>
          <ListItemButton>
            <ListItemText primary={props.text} />
          </ListItemButton>
        </ListItem>
      </Link>
    </>
  );
};

export default SideNavigationItem;
