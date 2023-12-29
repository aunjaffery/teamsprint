import { Box, keyframes } from "@chakra-ui/react";
import { motion } from "framer-motion";
import LogoIcon from "../../assets/LogoIcon";

const LogoAnimation = ({ size = "60" }) => {
  const animationKeyframes = keyframes`
  20%,50%,80%,to { transform: translateY(0);}
  40%% { transform: translateY(-30px);}
  70% { transform: translateY(-15px);}
  90% { transform: translateY(-4px); }
`;
  const animation = `${animationKeyframes} 1s ease-in-out infinite`;
  return (
    <div>
      <Box as={motion.div} animation={animation}>
        <LogoIcon size={size} />
      </Box>
    </div>
  );
};

export default LogoAnimation;
