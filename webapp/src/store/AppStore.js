import { Drawer_Links } from "../services/Drawer_Links";

const createAppStore = (set, get) => ({
  isSidebarOpen: false,
  sidebarLinks: Drawer_Links,
  onSidebarOpen: () => set(() => ({ isSidebarOpen: true })),
  onSidebarClose: () => {
    let list = get().sidebarLinks;
    list.map((x) => {
      if (x.open) {
        x.open = false;
      }
      return x;
    });
    set({ isSidebarOpen: false, sidebarLinks: [...list] });
  },
  onOpenSubMenu: (props) => {
    let list = get().sidebarLinks;
    let l = list.find((x) => x.id === props);
    if (l.open) {
      l.open = false;
    } else {
      l.open = true;
    }
    set({ sidebarLinks: [...list] });
  },
});
export default createAppStore;
