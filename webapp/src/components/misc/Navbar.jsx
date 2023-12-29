import { Avatar, Box, Flex, IconButton, Text } from "@chakra-ui/react";
import { AiOutlineMenuFold, AiOutlineMenuUnfold } from "react-icons/ai";
import useBoundStore from "../../store/Store";
import { MdMenu, MdNotifications } from "react-icons/md";

const Navbar = () => {
  const { isSidebarOpen, onSidebarClose, onSidebarOpen } = useBoundStore(
    (state) => state,
  );
  return (
    <Box w="100%">
      <Flex
        bg="none"
        h="68px"
        align="center"
        mx="4"
        justify="space-between"
        borderBottom="1px"
        borderColor="gray.300"
      >
        <IconButton
          css={{
            WebkitTapHighlightColor: "rgba(0, 0, 0, 0)",
          }}
          size="xs"
          aria-label="close-btn"
          icon={<MdMenu size="26" />}
          onClick={isSidebarOpen ? onSidebarClose : onSidebarOpen}
          _hover={{ outline: "none" }}
          _focus={{ outline: "none" }}
          _active={{ bg: "none", outline: "none" }}
          color="black"
        />

        <Flex justifyContent="center" alignItems="center" pr="2">
          <Flex justifyContent="center" alignItems="center" pr="4">
            <MdNotifications size="24" />
          </Flex>
          <Flex justifyContent="center" alignItems="center" pr="2">
            <Avatar name="Dan Abrahmov" src="https://bit.ly/prosper-baba" />
            <Text pl="2" fontSize="sm" fontWeight="bold">
              AunJaffery
            </Text>
          </Flex>
        </Flex>
      </Flex>
    </Box>
  );
};

export default Navbar;
