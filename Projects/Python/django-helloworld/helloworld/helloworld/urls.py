from django.http import HttpResponse
from django.urls import path

# Define a simple view that returns "Hello, World!"
def hello_world(request):
    return HttpResponse("Hello, World!")

# URL patterns
urlpatterns = [
    path('', hello_world),
]
