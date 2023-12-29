import { Box, Flex, Image, Text } from "@chakra-ui/react";
import PageTitle from "../../components/misc/PageTitle";

const Dashboard = () => {
  return (
    <Box>
      <PageTitle title="Dasboard" />
      <Box>
        {[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map((x) => (
          <Box key={x} boxShadow="md">
            <Box bg="blackAlpha.50" mb="10" p="6" borderRadius="lg">
              <Text fontWeight="bold" fontSize="lg" mb="2">
                Lorem ipsum dolor sit amet, qui minim labore adipisicing minim
                sint cillum sint consectetur cupidatat.
              </Text>
              <Text color="gray.600" fontSize="sm">
                Lorem ipsum dolor sit amet, officia excepteur ex fugiat
                reprehenderit enim labore culpa sint ad nisi Lorem pariatur
                mollit ex esse exercitation amet. Nisi anim cupidatat excepteur
                officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip
                amet voluptate voluptate dolor minim nulla est proident. Nostrud
                officia pariatur ut officia. Sit irure elit esse ea nulla sunt
                ex occaecat reprehenderit commodo officia dolor Lorem duis
                laboris cupidatat officia voluptate. Culpa proident adipisicing
                id nulla nisi laboris ex in Lorem sunt duis officia eiusmod.
                Aliqua reprehenderit commodo ex non excepteur duis sunt velit
                enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur
                et est culpa et culpa duis.
              </Text>
            </Box>
          </Box>
        ))}
      </Box>
    </Box>
  );
};

export default Dashboard;
