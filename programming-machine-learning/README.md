python3 -m venv pml

source pml/bin/activate

deactivate

pip install -r requirements.txt

pip freeze > requirements.txt

pip install pandas seaborn pytest wcwidth