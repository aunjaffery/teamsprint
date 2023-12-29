import { useEffect } from "react";
import useBoundStore from "../../store/Store";
import jwtDecode from "jwt-decode";
import { setSession, getAccessToken } from "../../services/jwt.service";
import { SiFirefox } from "react-icons/si";
import { Box, Flex, Spinner } from "@chakra-ui/react";
import LogoAnimation from "../misc/LogoAnimation";

const Auth = ({ children }) => {
  const { loginWithToken, tokenLoading, logoutService } = useBoundStore(
    (state) => state,
  );

  const handleAuthentication = async () => {
    let access_token = getAccessToken();
    if (!access_token) {
      logoutService();
      return;
    }
    if (!isAuthTokenValid(access_token)) return;
    setSession(access_token);
    loginWithToken();
  };

  const isAuthTokenValid = (access_token) => {
    const decoded = jwtDecode(access_token);
    const currentTime = Date.now() / 1000;
    if (decoded.exp < currentTime) {
      console.warn("access token expired");
      logoutService();
      return false;
    } else {
      return true;
    }
  };
  useEffect(() => {
    handleAuthentication();
  }, []);

  return (
    <div>
      {tokenLoading ? (
        <Box bg="bg.100" h="100vh">
          <Flex justifyContent="center" alignItems="center" h="100%">
            <LogoAnimation />
          </Flex>
        </Box>
      ) : (
        children
      )}
    </div>
  );
};

export default Auth;
