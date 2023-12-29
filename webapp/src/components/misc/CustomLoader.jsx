import { Flex } from "@chakra-ui/react";
import LogoAnimation from "./LogoAnimation";

const CustomLoader = ({ height = "600px" }) => {
  return (
    <Flex justifyContent="center" alignItems="center" h={height}>
      <LogoAnimation size="60" />
    </Flex>
  );
};

export default CustomLoader;
