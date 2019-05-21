package words.fxui;

import javafx.fxml.FXML;
import javafx.scene.control.Label;
import javafx.scene.text.Text;
import javafx.stage.FileChooser;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;

public class Controller {

    @FXML
    private Label fileName;

    @FXML
    private Text output;

    @FXML
    public void chooseFile() {
        FileChooser fileChooser = new FileChooser();
        fileChooser.setTitle("Open Resource File");
        fileChooser.getExtensionFilters().addAll(
            new FileChooser.ExtensionFilter("Text Files", "*.txt"),
            new FileChooser.ExtensionFilter("All Files", "*.*"));

        File selectedFile = fileChooser.showOpenDialog(null);
        if (selectedFile != null) {
            fileName.setText(selectedFile.getName());

            Path path = selectedFile.toPath();
            try {
                List<String> strings = Files.readAllLines(path);
                output.setText(String.join(" ", strings));
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
