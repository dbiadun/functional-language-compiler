import "List.hs";
import "Either.hs";
import "Maybe.hs";

shiftL :: Int -> Int -> Int;
shiftR :: Int -> Int -> Int;

return :: a -> IO a;
putStr :: String -> IO ();
putInt :: Int -> IO ();
getLine :: IO String;

stateSetInt :: String -> Int -> IO ();
stateGetInt :: String -> IO Int;

-- Go functions
tinySleep :: Int -> IO ();

-- TinyGo microcontroller functions
tinyLED :: IO Int;
tinyLED2 :: IO Int;
tinyBUTTON :: IO Int;

tinyConfigure :: Int -> Int -> IO ();
tinyPinInput :: IO Int;
tinyPinInputPullup :: IO Int;
tinyPinInputPulldown :: IO Int;
tinyPinOutput :: IO Int;

tinyPinRising :: IO Int;
tinyPinFalling :: IO Int;
tinyPinToggle :: IO Int;

tinyGet :: Int -> IO Bool;
tinyLow :: Int -> IO ();
tinyHigh :: Int -> IO ();
tinySetInterrupt :: Int -> Int -> (Int -> IO ()) -> IO ();

tinySetTimer :: Int -> Int -> IO () -> IO ();
tinyStopTimer :: Int -> IO ();
tinyStartTimer :: Int -> IO ();

tinyI2CConfigure :: Int -> Int -> Int -> Int -> IO (Maybe String);
tinyI2CConfigureDefault :: Int -> IO (Maybe String);
tinyI2CReadRegister :: Int -> Int -> Int -> Int -> IO (Either String (List Int));
tinyI2CWriteRegister :: Int -> Int -> Int -> List Int -> IO (Maybe String);
tinyI2CTx :: Int -> Int -> List Int -> Int -> IO (Either String (List Int));