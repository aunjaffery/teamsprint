import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import useBoundStore from "../../store/Store";
import { Box, Button, Flex, Input, Text } from "@chakra-ui/react";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import Domain from "../../services/Endpoint";
import { toast } from "react-toastify";
import LogoIcon from "../../assets/LogoIcon";

const LoginPage = () => {
  const navigate = useNavigate();
  const { loginService, user } = useBoundStore((state) => state);
  const login = useMutation({
    mutationFn: async (e) => {
      e.preventDefault();
      const rsp = await axios.post(
        `${Domain}/user/login`,
        new FormData(e.target),
      );
      return rsp.data;
    },
    onSuccess: (data) => loginService(data.user, data.token),
    onError: () => toast.error("Error! Invalid credentials"),
  });
  useEffect(() => {
    if (!!user) {
      navigate("/");
    }
  }, [user]);

  return (
    <Box bg="bg.100" w="100%" h="100vh">
      <Flex justifyContent="center" alignItems="center" w="100%" h="100%">
        <Box bg="white" borderRadius="lg" boxShadow="md" minW="350px" py="6">
          <Box p="14">
            <Flex justifyContent="center" mb="6">
              <Text fontSize="3xl" fontWeight="bold">
                Welcome
              </Text>
            </Flex>
            <Flex justifyContent="center" mb="12">
              <Box>
                <LogoIcon size="80" color="white" />
              </Box>
            </Flex>
            <form onSubmit={login.mutate}>
              <Flex justifyContent="center" mb="8">
                <Input
                  minW="280px"
                  type="text"
                  name="email"
                  placeholder="Email"
                  required
                />
              </Flex>
              <Flex justifyContent="center" mb="8">
                <Input
                  minW="280px"
                  type="password"
                  name="password"
                  placeholder="Password"
                  required
                />
              </Flex>
              <Flex justifyContent="center" mb="8">
                <Button
                  w="full"
                  bgGradient="linear(to-r,#21d4fd,#b721ff)"
                  color="white"
                  isLoading={login.isPending}
                  loadingText="Logging in"
                  _hover={{
                    bgGradient: "linear(to-r,#21d4fd,#b721ff)",
                  }}
                  type="submit"
                >
                  Login
                </Button>
              </Flex>
            </form>
          </Box>
        </Box>
      </Flex>
    </Box>
  );
};

export default LoginPage;
