import { Box, Flex, Text } from "@chakra-ui/react";

const NotFound = () => {
  return (
    <Box h="600px">
      <Flex justifyContent="center" alignItems="center" h="100%">
        <Text fontWeight="bold" fontSize="4xl" color="gray.500">
          Page Not Found!
        </Text>
      </Flex>
    </Box>
  );
};

export default NotFound;
