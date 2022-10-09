import { ListItem, ListItemButton, ListItemText } from "@mui/material";

interface Props {
  text: string;
}

const SideNavigationItem: React.FC<Props> = (props: Props) => {
  return (
    <>
      <ListItem>
        <ListItemButton>
          <ListItemText primary={props.text} />
        </ListItemButton>
      </ListItem>
    </>
  );
};

export default SideNavigationItem;
