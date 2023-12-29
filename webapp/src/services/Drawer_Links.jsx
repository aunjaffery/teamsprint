import { AiFillHome, AiFillSetting, AiFillTag } from "react-icons/ai";
import { ImStatsDots } from "react-icons/im";
import { MdInventory } from "react-icons/md";
import { BsFillPeopleFill } from "react-icons/bs";
import { FaUserCheck } from "react-icons/fa";

export const Drawer_Links = [
  {
    id: 1,
    name: "Dashboard",
    path: "/",
    icon: <AiFillHome size="20" />,
    child_routes: null,
  },
  {
    id: 2,
    name: "Products",
    path: "/products",
    icon: <AiFillTag size="20" />,
    child_routes: [
      {
        id: 1,
        name: "Add product",
        path: "/addproduct",
      },
      {
        id: 2,
        name: "List products",
        path: "/listproducts",
      },
      {
        id: 3,
        name: "Brands",
        path: "/brands",
      },
    ],
  },
  {
    id: 3,
    name: "Workspace",
    path: "/ws",
    icon: <MdInventory size="20" />,
    child_routes: null,
  },
  {
    id: 4,
    name: "Kanban",
    path: "/kanban",
    icon: <BsFillPeopleFill size="20" />,
    child_routes: null,
  },
  {
    id: 5,
    name: "Reporting",
    path: "/reporting",
    icon: <ImStatsDots size="20" />,
    child_routes: null,
  },
  {
    id: 6,
    name: "Users",
    path: "/users",
    icon: <FaUserCheck size="20" />,
    child_routes: null,
  },
  {
    id: 7,
    name: "Settings",
    path: "/settings",
    icon: <AiFillSetting size="20" />,
    child_routes: null,
  },
];
