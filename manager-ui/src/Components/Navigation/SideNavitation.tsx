import { Box, List, ListItemIcon, ListSubheader } from "@mui/material";
import React, { useEffect, useState } from "react";
import SideNavigationItem from "./SideNavigationItem";
import { menus, Menu, Link as LinkInterface } from "./Helpers/Routes";

const backColor = "lightgrey";

const SideNavigation: React.FC = () => {
  const [sideNavigationMenu, setSideNavigationMenu] = useState(menus);

  useEffect(() => {
    let newNavigation = [...sideNavigationMenu];
    newNavigation[0].links.push({ name: "hello", to: "world" }); // Api call goes here
    setSideNavigationMenu(newNavigation);
  }, []);

  return (
    <div>
      <Box
        style={{
          width: "300px",
          height: "100vh",
          boxShadow: "5px 4px 10px 5px rgba(0,0,0,0.10)",
          backgroundColor: backColor,
        }}
      >
        {/* Loop through the menus */}
        {sideNavigationMenu.map((menu: Menu, index: number) => {
          return (
            <List
              key={index}
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
              {menu.links.map((link: LinkInterface, index: number) => (
                <SideNavigationItem key={index} text={link.name} to={link.to} />
              ))}
            </List>
          );
        })}
      </Box>
    </div>
  );
};

export default SideNavigation;
