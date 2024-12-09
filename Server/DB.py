import sqlite3

class Database:
    """
        Handles the database for the Watchtower Server

        Parameters:
            - filename (str) - Database filename
            - debug (bool) - Enable debug mode
         
        """
    def __init__(self, **kwargs):
        # Get kwargs from creating DB
        self.filename = kwargs.get('filename', 'WatchtowerData')
        self.debug = kwargs.get('debug', 'False')

        self.conn = sqlite3.connect(self.filename + ".db", check_same_thread=False)

        # Create tables for each section
        self.tables = {
            "system": """
                UserID INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
                Hostname TEXT,
                OS TEXT,
                Platform TEXT,
                PlatformVer TEXT,
                KernelVer TEXT,
                Architecture TEXT,
                Uptime TEXT,
                Boottime DATETIME,
                LoggedIn TEXT,
                LoggedTime DATETIME UNIQUE 
            """,
            "memory": """
                UserID INTEGER REFERENCES system(UserID),
                Total INTEGER,
                Available INTEGER,
                Used INTEGER,
                UsedPercent FLOAT,
                SwapTotal INTEGER,
                SwapUsed INTEGER,
                SwapUsedPercent FLOAT,
                LoggedTime DATETIME
            """,
            "cpu": """
                UserID INTEGER REFERENCES system(UserID),
                Cores INTEGER,
                UsagePerCore TEXT,
                AverageUsage FLOAT(64),
                ClockSpeedMhz FLOAT(64),
                LoggedTime DATETIME
            """,
            "network": """
                UserID INTEGER REFERENCES system(UserID),
                InterfaceName TEXT,
                IPAddresses TEXT,
                MACAddress TEXT,
                UploadSpeed FLOAT(64),
                DownloadSpeed FLOAT(64),
                LoggedTime DATETIME
            """,
            "disk": """
                UserID INTEGER REFERENCES system(UserID),
                Device TEXT,
                MountPoint TEXT,
                TotalSpace INTEGER,
                FreeSpace INTEGER,
                UsedSpace INTEGER,
                Usage FLOAT(64),
                LoggedTime DATETIME
            """,
            "processes": """
                UserID INTEGER REFERENCES system(UserID),
                Name TEXT,
                PID INTEGER(32),
                CPUPercent FLOAT(64),
                MemoryPercent FLOAT(32),
                LoggedTime DATETIME
            """
            }

        self.createTables()
        self.cursor = self.conn.cursor()

    def createTables(self) -> None:
        """
        Creates the required tables for the database

        Parameters:
                - None

            Returns:
                - None
        
        """
        self.cursor = self.conn.cursor()
        for table_name, schema in self.tables.items():
            query = f"CREATE TABLE IF NOT EXISTS {table_name} ({schema})"
            self.cursor.execute(query)
            if self.debug:
                print(f"Executed: {query}")

        # Commit the changes
        self.conn.commit()
        return

    def disp_rows(self, selectedTable: str) -> str:
        """
            Returns all the infomation in selected table - Used for debug

            Parameters:
                - selectedTable (str) - Table to view data from

            Returns:
                - Data (str) - Data currently in the DB 
        
        """
        Query = 'SELECT * FROM ' + selectedTable
        self.cursor.execute(Query)
        Output = ""
        for row in self.cursor:
            Output = Output + str(row)
        return Output

    def insert(self, selectedTable: str, columns: list, data: list) -> None:
        """

            Inserts data into database

            Parameters:
                - selectedTable (str) - Table to add data into
                - columns (list) - Columns to put the info into
                - data (list) - Data you want to insert into the database

            Returns:
                - None       
        """
        if len(columns) != len(data):
            print(f"Columns: {len(columns)} {columns} \n - Data: {len(data)} {data}")
            raise("Data cannot fit into columns")
        if not isinstance(columns, list) or not isinstance(data, list):
            raise("Data or columns must be list!")
        columns.append("LoggedTime")
        placeholders = ", ".join(["?"] * len(data))  
        query = f"INSERT INTO {selectedTable} ({', '.join(columns)}) VALUES ({placeholders}, CURRENT_TIMESTAMP)"
        
        # Data may be a list, convert to string so the DB can handle it
        data = [
        ", ".join(map(str, item)) if isinstance(item, list) else item
        for item in data
        ]

        if self.debug == True:
            print("Query:", query)
            print("Data:", data)
        self.cursor.execute(query, data)
        self.conn.commit()

    def check_userID(self, hostname: str) -> int:
        """
        Retrieves the UserID associated with a given Hostname.

        Parameters:
            - hostname (str): The Hostname to look up.

        Returns:
            - int: The UserID associated with the Hostname, or None if not found.
        """
        query = "SELECT UserID FROM system WHERE Hostname = ?"
        self.cursor.execute(query, (hostname,))
        result = self.cursor.fetchone()
        if result:
            if self.debug:
                print(f"Found UserID: {result[0]} for Hostname: {hostname}")
            return result[0]
        else:
            if self.debug:
                print(f"No UserID found for Hostname: {hostname}")
            return None

if __name__ == "__main__":
    db = Database(debug=True)
    #db.insert("system", ())
    db.insert("cpu", ["UserID", "Cores", "UsedPerCore", "AverageUsage", "ClockSpeedMhz"], [1, 8, 15.7, 12.3, 3200.0])
    print(db.disp_rows("cpu"))