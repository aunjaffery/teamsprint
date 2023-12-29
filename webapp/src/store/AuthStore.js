import Domain from "../services/Endpoint";
import axios from "axios";
import { setSession } from "../services/jwt.service";

const createAuthStore = (set) => ({
  user: null,
  authLoading: false,
  tokenLoading: true,
  setUser: (args) => set({ user: args }),
  logoutService: () => {
    setSession(null);
    set({ user: null, authLoading: false, tokenLoading: false });
  },
  loginService: (user, token) => {
    console.log(user);
    if (user && token) {
      setSession(token);
      set({ user: user });
    } else {
      set({ user: null });
    }
  },
  loginWithToken: async () => {
    try {
      const rsp = await axios(`${Domain}/user/validation`);
      if (rsp.data?.user && rsp.data?.token) {
        setSession(rsp.data?.token);
        set({ user: rsp.data?.user, tokenLoading: false });
      } else {
        set({ tokenLoading: false, user: null });
      }
    } catch (error) {
      console.log(error);
      logoutService();
    }
  },
});
export default createAuthStore;
