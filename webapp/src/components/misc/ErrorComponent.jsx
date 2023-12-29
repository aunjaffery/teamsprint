import { Box, Flex, Text } from "@chakra-ui/react";
const ErrorComponent = ({
  height = "600px",
  error = "Error! Somthing went Wrong",
}) => {
  return (
    <Box h={height}>
      <Flex justifyContent="center" alignItems="center" h="100%">
        <Text fontWeight="bold" fontSize="2xl" color="gray.500">
          {error}
        </Text>
      </Flex>
    </Box>
  );
};

export default ErrorComponent;
