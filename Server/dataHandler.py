
def handler(dbObject: object, data: str, debug: bool = False) -> None:
    """
        Handles the data into the database
        Parameters:
            - dbObject (object): The object of the database.
            - data (str): Data given from client
            - debug (bool): Enable debug mode

        Returns:
            - None
    """
        
    for section in data:
        if section == "system":
            for subsection in data[section]:
                    if debug == True:
                        print(f"{subsection}: {data[section][subsection]} - {type(data[section][subsection])}")
                    subsectionList.append(subsection)
                    subsectionDataList.append(data[section][subsection])  

            if (UserID := dbObject.check_userID(data[section]["Hostname"])):
                subsectionList.insert(0, "UserID")
                subsectionDataList.insert(0, UserID)
                dbObject.update(section, subsectionList, subsectionDataList, "Hostname = ?", [UserID])

            else:
                subsectionList.insert(0, "UserID")
                subsectionDataList.insert(0, "NULL")
                dbObject.insert(section, subsectionList, subsectionDataList)
        else:
            # Filter for dictonary datatype                
            if isinstance(data[section], dict):
                subsectionList = ["UserID"]
                subsectionDataList = ["2"]
                for subsection in data[section]:
                    if debug == True:
                        print(f"{subsection}: {data[section][subsection]} - {type(data[section][subsection])}")

                    subsectionList.append(subsection)
                    subsectionDataList.append(data[section][subsection])
                dbObject.insert(section, subsectionList, subsectionDataList)
            # Filter for list datatype
            elif isinstance(data[section], list):
                subsectionList = ["UserID"]
                subsectionDataList = ["2"]
                for item in data[section]:
                    for name in item:
                        subsectionList.append(name)
                        subsectionDataList.append(item[name])
                    dbObject.insert(section, subsectionList, subsectionDataList)
                    subsectionList = ["UserID"]
                    subsectionDataList = ["2"]
            # Filter for 'Other' datatype
            else:
                return NotImplemented
    
    return None


if __name__ == '__main__':
    handler()