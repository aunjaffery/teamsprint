import create from "zustand";
import { devtools } from "zustand/middleware";
import AuthStore from "./AuthStore";
import AppStore from "./AppStore";

const useBoundStore = create()(
  devtools((...a) => ({
    ...AuthStore(...a),
    ...AppStore(...a),
  })),
);
export default useBoundStore;
