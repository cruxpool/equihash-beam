cmake_minimum_required(VERSION 3.5.1)

set(CMAKE_CXX_STANDARD 14)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -pthread")
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

file(GLOB hasher_beamhash_SRC
        "src/*/*.c"
        "src/*.c"
        "src/*.cpp"
        "src/*/*.cpp"
        )


add_library(beam_shared SHARED ${hasher_beamhash_SRC})

set_target_properties(beam_shared PROPERTIES OUTPUT_NAME "beampow")

string(REPLACE "\\" "/" PROJECT_SOURCE_DIR ${PROJECT_SOURCE_DIR})
    message(STATUS "include_directories ${PROJECT_SOURCE_DIR}")
    include_directories(${PROJECT_NAME} PUBLIC ${PROJECT_SOURCE_DIR})
    include_directories(${PROJECT_NAME} PUBLIC "./src")

    set_target_properties(beam_shared PROPERTIES OUTPUT_NAME "beampow")

set(LIBRARY_OUTPUT_PATH ${PROJECT_BINARY_DIR}/lib)

install(TARGETS beam_shared
	LIBRARY DESTINATION lib
	ARCHIVE DESTINATION lib)
