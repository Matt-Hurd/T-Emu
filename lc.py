import os

def count_lines_in_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        return len(file.readlines())

def count_lines_in_directory(directory):
    total_lines = 0
    for root, _, files in os.walk(directory):
        for file in files:
            if file.endswith('.go'):
                file_path = os.path.join(root, file)
                total_lines += count_lines_in_file(file_path)
    return total_lines

if __name__ == "__main__":
    directory_path = input("Enter the directory path: ")
    if os.path.isdir(directory_path):
        total_lines = count_lines_in_directory(directory_path)
        print(f"Total lines in .go files: {total_lines}")
    else:
        print("The provided path is not a valid directory.")
