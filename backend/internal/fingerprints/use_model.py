from tensorflow import keras
from skimage.io import imread
from skimage.transform import resize
import numpy as np

# Load the trained model
siamese_model = keras.models.load_model('fingerprint_check_model.keras')

# Load and preprocess the images
image1 = imread('c:/input prints/4_1.jpg', as_gray=True)
image1 = resize(image1, (328, 356)) / 255.0

image2 = imread('c:/input prints/4_1.jpg', as_gray=True)
image2 = resize(image2, (328, 356)) / 255.0

# Prepare the images as pairs
pair = np.array([[image1, image2]])

# Predict the similarity score
similarity_score = siamese_model.predict([pair[:, 0], pair[:, 1]])

if similarity_score > 0.5:
    print("This is the same person")
else:
    print("This is a different person")

print(similarity_score)
