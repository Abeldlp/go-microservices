import InboxIcon from "@mui/icons-material/Shop";
import SettingsIcon from "@mui/icons-material/Settings";
import { api } from "../../../api/Api";

interface Menu {
  title: string;
  subroute: string;
  icon: any;
  links: Link[];
}

interface Link {
  name: string;
  to: string;
}

api.get("/my-api-call");

const menus: Menu[] = [
  {
    title: "Shops",
    subroute: "/shops",
    icon: InboxIcon,
    links: [
      {
        name: "Murasaki sports",
        to: "/123",
      },
      {
        name: "Rookies malpan",
        to: "/124",
      },
    ],
  },
  {
    title: "Settings",
    subroute: "/settings",
    icon: SettingsIcon,
    links: [
      {
        name: "admin",
        to: "/124",
      },
    ],
  },
];

export { menus, type Menu, type Link };
