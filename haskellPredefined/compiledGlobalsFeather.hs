import "compiledGlobals.hs";
import "List.hs";
import "Either.hs";
import "Maybe.hs";

-------------------------------------- GPIO PINS ----------------------------------------

-- Returns a pin number corresponding to the first LED at the board
tinyLED :: IO Int;

-- Returns a pin number corresponding to the second LED at the board
tinyLED2 :: IO Int;

-- Returns a pin nummber corresponding to the button at the board
tinyBUTTON :: IO Int;

-------------------------------- GPIO PIN CONFIGURATION ---------------------------------

{- Configures a GPIO pin
-- 1. param - pin number
-- 2. param - number indicating pin mode
-}
tinyConfigure :: Int -> Int -> IO ();

-- Input pin mode
tinyPinInput :: IO Int;

-- Input pin mode with pull-up resistor
tinyPinInputPullup :: IO Int;

-- Input pin mode with pull-down resistor
tinyPinInputPulldown :: IO Int;

-- Output pin mode
tinyPinOutput :: IO Int;

----------------------------------- GPIO PIN STATE --------------------------------------

{- Returns current state on the pin
-- 1. param - pin number
-}
tinyGet :: Int -> IO Bool;

{- Sets a high state on the pin
-- 1. param - pin number
-}
tinyLow :: Int -> IO ();

{- Sets a high state on the pin
-- 1. param - pin number
-}
tinyHigh :: Int -> IO ();

----------------------------------- GPIO INTERRUPTS -------------------------------------

{- Sets an interrupt on a GPIO pin
-- 1. param - pin number
-- 2. param - number indicating an event triggering the interrupt
-- 3. parm - function called when interrupt occurs (the only param is a pin number (doesn't have to be the same
-- one that we set interrupt on))
-}
tinySetInterrupt :: Int -> Int -> (Int -> IO ()) -> IO ();

-- Number indicating pin rising event
tinyPinRising :: IO Int;

-- Number indicating pin falling event
tinyPinFalling :: IO Int;

-- Number indicating pin rising or falling event
tinyPinToggle :: IO Int;

--------------------------------------- TIMERS ------------------------------------------

{- Configures and sets timer with an interrupt
-- 1. param - timer number (from 0 to 4)
-- 2. param - time space between ticks in microseconds (32 bit value)
-- 3. param - function called when interrupt occurs (on every tick)
-}
tinySetTimer :: Int -> Int -> IO () -> IO ();

{- Stops timer
-- 1. param - timer number
-}
tinyStopTimer :: Int -> IO ();

{- Continue counting of the stopped timer
-- 1. param - timer number
-}
tinyStartTimer :: Int -> IO ();

---------------------------------------- I2C --------------------------------------------

{- Configures I2C
-- 1. param - number indicating an I2C bus (0 or 1)
-- 2. param - frequency
-- 3. param - SCL pin number
-- 4. param - SDA pin number
-- returns Just (with error message) if error occurs and Nothing when executed without issues
-}
tinyI2CConfigure :: Int -> Int -> Int -> Int -> IO (Maybe String);

{- Configures I2C with the default configuration
-- 1. param - I2C bus number
-- returns Just (with error message) if error occurs and Nothing when executed without issues
-}
tinyI2CConfigureDefault :: Int -> IO (Maybe String);

{- Reads register from I2C
-- 1. param - I2C bus number
-- 2. param - device address
-- 3. param - register address
-- 4. param - number of bytes to read
-- returns Left (with error message) if error occurs or Right (with list of bytes read) if execution ends normally
-}
tinyI2CReadRegister :: Int -> Int -> Int -> Int -> IO (Either String (List Int));

{- Writes to register with I2C
-- 1. param - I2C bus number
-- 2. param - device address
-- 3. param - register address
-- 4. param - list of bytes to write
-- returns Just (with error message) if error occurs and Nothing when execution ends normally
-}
tinyI2CWriteRegister :: Int -> Int -> Int -> List Int -> IO (Maybe String);

{- Does a single I2C transaction
-- 1. param - I2C bus number
-- 2. param - device address
-- 3. param - list of bytes to write
-- 4. param - number of bytes to read after write
-- returns Left (with error message) if error occurs on Right (with list of bytes read) if execution ends normally
-}
tinyI2CTx :: Int -> Int -> List Int -> Int -> IO (Either String (List Int));
