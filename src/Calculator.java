import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Calculator extends JFrame implements ActionListener {
    private JTextField display;
    private String operator = "";
    private double num1 = 0;
    private double num2 = 0;
    private double result = 0;

    public static void main(String[] args) {
        new Calculator();
    }

    public Calculator() {

        setTitle("Calculator");
        setSize(300, 400);
        setDefaultCloseOperation(EXIT_ON_CLOSE);
        setLayout(new BorderLayout());


        display = new JTextField();
        display.setEditable(false);
        display.setFont(new Font("Arial", Font.BOLD, 24));
        add(display, BorderLayout.NORTH);

        JPanel buttonPanel = new JPanel();
        buttonPanel.setLayout(new GridLayout(4, 4, 10, 10));

        String[] buttons = {
                "7", "8", "9", "/",
                "4", "5", "6", "*",
                "1", "2", "3", "-",
                "C", "0", "=", "+"
        };

        for (String text : buttons) {
            JButton button = new JButton(text);
            button.setFont(new Font("Arial", Font.BOLD, 20));
            button.addActionListener(this);
            buttonPanel.add(button);
        }

        add(buttonPanel, BorderLayout.CENTER);

        setVisible(true);
    }

    @Override
    public void actionPerformed(ActionEvent e) {
        String command = e.getActionCommand();

        if (command.matches("\\d")) { // If it's a digit
            display.setText(display.getText() + command);
        } else if (command.matches("[+\\-*/]")) { // If it's an operator
            num1 = Double.parseDouble(display.getText());
            operator = command;
            display.setText("");
        } else if (command.equals("=")) { // If it's equals
            num2 = Double.parseDouble(display.getText());
            switch (operator) {
                case "+" -> result = num1 + num2;
                case "-" -> result = num1 - num2;
                case "*" -> result = num1 * num2;
                case "/" -> result = num1 / num2;
            }
            display.setText(String.valueOf(result));
        } else if (command.equals("C")) { // Clear the display
            display.setText("");
            num1 = num2 = result = 0;
            operator = "";
        }
    }
}
