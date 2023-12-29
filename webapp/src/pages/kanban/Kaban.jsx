import { Box, Button, Flex, Text } from "@chakra-ui/react";
import PageTitle from "../../components/misc/PageTitle";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import Domain from "../../services/Endpoint";
import ErrorComponent from "../../components/misc/ErrorComponent";
import CustomLoader from "../../components/misc/CustomLoader";
import { LuSettings2, LuUsers } from "react-icons/lu";
import { IoMdAdd, IoMdAddCircle } from "react-icons/io";

const Kaban = () => {
  const { isError, data, isLoading, isFetching } = useQuery({
    queryKey: ["workspacekbns"],
    queryFn: () =>
      axios.get(`${Domain}/ws/workspacekbns`).then((res) => res.data),
  });
  console.log(isError, isLoading, isFetching);
  console.log(data);
  if (isError) {
    return <ErrorComponent error="Error! Please try again." />;
  }
  const ws = data?.ws;
  // const ws = [
  //   {
  //     id: "658c7c2109562acd997f6e71",
  //     name: "aunox awsome",
  //     creator: "658c44dcecce4ed04123645b",
  //     members: ["658c44dcecce4ed04123645b", "658c407badd39cee5d75f6ce"],
  //     kanban: [
  //       {
  //         id: "658d6b52c22ad481c0a53b62",
  //         title: "new kanban",
  //         creator: "658c44dcecce4ed04123645b",
  //         visibility: "workspace",
  //       },
  //     ],
  //   },
  // ];
  // const isLoading = false;
  return (
    <Box>
      <PageTitle title="Your Workspaces" />
      <Box>
        {isLoading ? (
          <CustomLoader />
        ) : (
          <Flex direction="column" gridRowGap="12" mt="10">
            {ws &&
              ws.map((w) => (
                <Box key={w.id} w="full">
                  <Flex
                    justifyContent="space-between"
                    alignItems="center"
                    w="full"
                    maxW="800px"
                  >
                    <Flex alignItems="center" gridColumnGap="4">
                      <Placeholder w={w.name} />
                      <Text key={w.id} fontWeight="bold" fontSize="xl">
                        {w.name}
                      </Text>
                    </Flex>
                    <Flex alignItems="center" gridColumnGap="4">
                      <Button
                        leftIcon={<LuUsers />}
                        bg="blackAlpha.200"
                        size="sm"
                      >
                        Members (3)
                      </Button>
                      <Button
                        leftIcon={<LuSettings2 />}
                        bg="blackAlpha.200"
                        size="sm"
                      >
                        Settings
                      </Button>
                    </Flex>
                  </Flex>
                  <Box mt="6">
                    <Flex
                      justifyContent="flex-start"
                      alignItems="center"
                      gridColumnGap="6"
                    >
                      {w.kanban &&
                        w.kanban.map((k) => (
                          <Box
                            key={k.id}
                            w="240px"
                            h="120px"
                            bg="gray.300"
                            borderRadius="lg"
                            bgGradient="linear(to-r, #3d2cbd, #dcbdf7)"
                          >
                            <Text
                              px="4"
                              pt="4"
                              fontSize="lg"
                              fontWeight="bold"
                              color="white"
                              noOfLines={1}
                            >
                              {k.title}
                            </Text>
                          </Box>
                        ))}
                      <Flex
                        w="240px"
                        h="120px"
                        bg="gray.300"
                        borderRadius="lg"
                        bgGradient="linear(to-r, gray.300, gray.200)"
                        direction="column"
                        justifyContent="center"
                        alignItems="center"
                      >
                        <Text
                          pt="4"
                          color="gray.600"
                          noOfLines={1}
                          textAlign="center"
                          fontSize="sm"
                        >
                          Create new board
                        </Text>
                        <Flex justifyContent="center" mt="1" color="gray.500">
                          <IoMdAddCircle size="24" />
                        </Flex>
                      </Flex>
                    </Flex>
                  </Box>
                </Box>
              ))}
          </Flex>
        )}
      </Box>
    </Box>
  );
};
const Placeholder = ({ w = "v", wh = "40px" }) => {
  let alpha = w.charAt(0).toUpperCase();
  return (
    <Flex
      w={wh}
      h={wh}
      bg="black"
      borderRadius="md"
      bgGradient="linear(to-r, teal.500, green.500)"
      justifyContent="center"
      alignItems="center"
    >
      <Text color="white" fontWeight="bold" fontSize="xl">
        {alpha}
      </Text>
    </Flex>
  );
};

export default Kaban;
