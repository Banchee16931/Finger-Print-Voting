import os
import numpy as np
import tensorflow as tf
from tensorflow import keras
from keras import layers
from sklearn.model_selection import train_test_split
from skimage.io import imread
from skimage.transform import resize
from tqdm import tqdm
import random
from PIL import Image

def grayscale_jpg_to_array(image_path):
    # Open the image
    image = Image.open(image_path).convert("L")  # Convert to grayscale

    # Convert the PIL image to a NumPy array
    image_array = np.array(image)

    return image_array

# Load and preprocess the dataset
def load_dataset(folder_path):
    image_files = [f for f in os.listdir(folder_path) if f.endswith(".jpg")]
    images = []
    labels = []
    
    for file in image_files:
        user_id, scan_num = map(int, os.path.splitext(file)[0].split("_"))
        image = grayscale_jpg_to_array(os.path.join(folder_path, file))
        images.append(image)
        labels.append(user_id)
    
    images = np.array(images)
    labels = np.array(labels)

    return images, labels

# Create a Siamese network model
def create_siamese_model(input_shape):
    base_model = keras.Sequential([
        layers.Reshape((input_shape[0], input_shape[1], 1), input_shape=input_shape),
        layers.Conv2D(32, (3, 3), activation='relu'),
        layers.MaxPooling2D((2, 2)),
        layers.Conv2D(64, (3, 3), activation='relu'),
        layers.MaxPooling2D((2, 2)),
        layers.Flatten(),
        layers.Dense(128, activation='relu'),
    ])
    
    input_a = keras.Input(shape=input_shape)
    input_b = keras.Input(shape=input_shape)
    
    features_a = base_model(input_a)
    features_b = base_model(input_b)
    
    distance = tf.norm(features_a - features_b, axis=1)
    distance = keras.layers.Reshape((1,))(distance)
    
    output = layers.Dense(1, activation='tanh')(distance)
    
    siamese_model = keras.Model(inputs=[input_a, input_b], outputs=output)
    return siamese_model

def generate_matching_pairs(image_list, label_list, batch_size):
    matching_pairs = []

    while len(matching_pairs) < batch_size:
        index1, index2 = random.sample(range(len(image_list)), 2)

        label1 = label_list[index1]
        label2 = label_list[index2]

        if label1 == label2:
            matching_pairs.append([image_list[index1], image_list[index2]])

    return matching_pairs

def generate_non_matching_pairs(image_list, label_list, batch_size):
    non_matching_pairs = []

    while len(non_matching_pairs) < batch_size:
        index1, index2 = random.sample(range(len(image_list)), 2)

        label1 = label_list[index1]
        label2 = label_list[index2]

        if label1 != label2:
            non_matching_pairs.append([image_list[index1], image_list[index2]])

    return non_matching_pairs

def prepare_pairs(image_list, label_list, batch_size):
    matching_pairs = generate_matching_pairs(image_list, label_list, batch_size * 50)
    non_matching_pairs = generate_non_matching_pairs(image_list, label_list, batch_size * 50)

    combined_pairs = matching_pairs + non_matching_pairs

    labels = [1] * len(matching_pairs) + [0] * len(non_matching_pairs)

    # Combine the two lists into pairs using zip
    combined = list(zip(combined_pairs, labels))

    # Shuffle the combined list
    random.shuffle(combined)

    # Unzip the shuffled pairs back into separate lists
    shuffled_combined_pairs, shuffled_labels = zip(*combined)

    #print(shuffled_combined_pairs)

    return np.array(shuffled_combined_pairs), np.array(shuffled_labels)

def main():
    folder_path = "C:/input prints"
    images, labels = load_dataset(folder_path)
    
    train_images, val_images, train_labels, val_labels = train_test_split(images, labels, test_size=0.2, random_state=42)
    
    input_shape = images.shape[1:]
    
    siamese_model = create_siamese_model(input_shape)
    siamese_model.compile(optimizer='adam', loss='binary_crossentropy', metrics=['accuracy'])
    
    batch_size = 16  # Adjust this based on available memory
    train_pairs, train_labels = prepare_pairs(train_images, train_labels, batch_size=batch_size)

    val_pairs, val_labels = prepare_pairs(val_images, val_labels, batch_size=batch_size)
    
    siamese_model.fit([train_pairs[:, 0], train_pairs[:, 1]], train_labels,
                      validation_data=([val_pairs[:, 0], val_pairs[:, 1]], val_labels),
                      batch_size=batch_size, epochs=100)
    
    siamese_model.save('fingerprint_check_model.keras')

if __name__ == "__main__":
    main()
