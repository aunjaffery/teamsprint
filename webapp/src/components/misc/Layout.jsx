import { useLocation } from "react-router-dom";
import Navbar from "./Navbar";
import { Box, Container, Flex, useMediaQuery } from "@chakra-ui/react";
import useBoundStore from "../../store/Store";
import { useEffect } from "react";
import DrawerComp from "../misc/Drawer";

const Layout = ({ children }) => {
  const location = useLocation();
  const [isLarger] = useMediaQuery("(min-width: 768px)");
  const { isSidebarOpen, onSidebarOpen } = useBoundStore((state) => state);

  const render = () => {
    return location.pathname === "/login" || location.pathname === "/signup"
      ? true
      : false;
  };
  useEffect(() => {
    if (isLarger) {
      onSidebarOpen();
    } else return;
  }, []);

  if (render()) {
    return <Box>{children}</Box>;
  }

  return (
    <Box position="relative">
      <DrawerOverLay isLarger={isLarger} isSidebarOpen={isSidebarOpen} />
      <Flex>
        <DrawerComp isLarger={isLarger} />
        <Box
          flex="1"
          minH="calc(100vh)"
          bg="bg.100"
          position="relative"
          overflowY="hidden"
        >
          <Box
            h="100vh"
            overflowY="auto"
            sx={{
              "&::-webkit-scrollbar-track": {
                bg: "transparent",
              },
              "&::-webkit-scrollbar": {
                width: "4px",
              },
              "&::-webkit-scrollbar-thumb": {
                bg: "blackAlpha.400",
                borderRadius: "20px",
              },
            }}
          >
            <Navbar />
            <Container maxW="container.xl">{children}</Container>
          </Box>
        </Box>
      </Flex>
    </Box>
  );
};
const DrawerOverLay = ({ isLarger, isSidebarOpen }) => {
  return (
    <Box>
      {!isLarger && isSidebarOpen && (
        <Box
          w="100vw"
          h="100vh"
          bg="blackAlpha.500"
          position="fixed"
          top="0"
          left="0"
          zIndex="1300"
        />
      )}
    </Box>
  );
};

export default Layout;
