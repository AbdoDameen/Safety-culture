# Safety Culture Internship Technical Assessment 2023
This repository contains my submission for the 2023 technical assessment for prospective SC interns. The assessment involved working with Go to demonstrate coding and testing skills.

# Folder Data Fetcher Module

## Overview
The `folders.go` module is designed to manage and retrieve folder data for different organizations.

### Functions

#### GetAllFolders
- **Purpose**: Retrieves a list of folders based on a given request.
- **Input**: Expects a `FetchFolderRequest` object containing criteria like `OrgID`.
- **Output**: Provides a `FetchFolderResponse` with folder details or an error if the operation fails.
- **Process**:
  - Initialises a slice to store folder data.
  - Calls `FetchAllFoldersByOrgID` using the `OrgID` from the request.
  - Transforms the fetched folders into a slice of pointers.
  - Returns a response filled with the folder pointers.

#### FetchAllFoldersByOrgID
- **Purpose**: Fetches folders belonging to a specific organization.
- **Input**: Requires an `orgID`.
- **Output**: Returns a slice of pointers to the relevant folders.
- **Process**:
  - Invokes `GetSampleData` to get folder data, which suggests usage of preset data, not live data.
  - Filters the folders to find those associated with the `orgID`.

### Observations and Improvements

- The `GetAllFolders` function can be optimized by removing redundant slices and directly working with folder pointers.
- Unused variables in loops should be replaced with `_` to indicate they are not used.
- Error handling needs to be addressed in `GetAllFolders` to manage potential issues during data fetching.
- If this module is meant for a production environment, `FetchAllFoldersByOrgID` should be modified to retrieve data from an actual database or live data source instead of static sample data.

---

By addressing these observations, the module can be more efficient and production-ready.
---

### Pagination Implementation Explanation:

In the updated GetAllFolders function, we've added two crucial parameters: 'limit' and 'offset'. These parameters are key to managing large datasets by breaking them down into manageable chunks. The 'limit' parameter specifies how many records to return in one go, while 'offset' determines where in the dataset to start returning records from.

Implementing pagination is essential for enhancing performance and user experience. It efficiently limits the data transferred over the network, reducing server load and leading to faster responses. Using 'offset' and 'limit' is a common practice as it's intuitive and scales well with large amounts of data.

Our API is now more flexible, allowing clients to specify exactly how much data they want to retrieve with each request. Additionally, the 'NextPage' flag in our responses makes it easy for clients to determine if there's more data to fetch. This approach not only optimizes performance but also brings our API in line with standard practices, making it more user-friendly and easier to integrate with.

This simple yet effective pagination strategy underscores our commitment to building scalable and efficient software that provides a smooth experience for users and developers alike.



