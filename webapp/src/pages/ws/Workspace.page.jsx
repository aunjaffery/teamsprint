import { Box, Flex, Text } from "@chakra-ui/react";
import Domain from "../../services/Endpoint";
import axios from "axios";
import { useQuery } from "@tanstack/react-query";
import PageTitle from "../../components/misc/PageTitle";
import CustomLoader from "../../components/misc/CustomLoader";
import ErrorComponent from "../../components/misc/ErrorComponent";

const Workspace = () => {
  const { error, data, isLoading } = useQuery({
    queryKey: ["findworkspace"],
    queryFn: () =>
      axios.get(`${Domain}/ws/findworkspace1`).then((res) => res.data),
  });
  const ws = data?.ws;
  if (error) {
    return <ErrorComponent height="600px" />;
  }
  return (
    <Box>
      <PageTitle title="Workspace" />
      <Text>This is Protected Route</Text>
      {isLoading ? (
        <CustomLoader />
      ) : (
        <Flex>
          {ws.map((w) => (
            <Text key={w.id}>{w.name}</Text>
          ))}
        </Flex>
      )}
    </Box>
  );
};

export default Workspace;
