import { Box, List, ListItemIcon, ListSubheader } from "@mui/material";
import React from "react";
import SideNavigationItem from "./SideNavigationItem";
import { menus, Menu, Link } from "./Helpers/Routes";

const backColor = "lightgrey";

const SideNavigation: React.FC = () => {
  return (
    <div>
      <Box
        style={{
          width: "400px",
          height: "100vh",
          boxShadow: "5px 4px 10px 5px rgba(0,0,0,0.10)",
          backgroundColor: backColor,
        }}
      >
        {/* Loop through the menus */}
        {menus.map((menu: Menu) => {
          return (
            <List
              component="nav"
              aria-labelledby="nested-list-subheader"
              subheader={
                <ListSubheader
                  component="div"
                  id="nested-list-subheader"
                  style={{
                    display: "flex",
                    alignItems: "center",
                    backgroundColor: backColor,
                  }}
                >
                  <ListItemIcon>{React.createElement(menu.icon)}</ListItemIcon>
                  <span>{menu.title}</span>
                </ListSubheader>
              }
            >
              {/* Loop through links in the menus */}
              {menu.links.map((link: Link) => (
                <SideNavigationItem text={link.name} />
              ))}
            </List>
          );
        })}
      </Box>
    </div>
  );
};

export default SideNavigation;
