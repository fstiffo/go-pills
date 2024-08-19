Building a software system to manage stock, a storeroom, or a shed involves careful planning and following best practices in software design. Here are some patterns and practices that can guide you in creating an efficient and scalable system:

### 1. **Domain-Driven Design (DDD)**
   - **Focus on the Domain**: Understand the core business logic and processes involved in inventory management.
   - **Bounded Context**: Define clear boundaries within your system to separate different areas of responsibility (e.g., inventory, orders, suppliers).
   - **Entities and Value Objects**: Use entities to represent objects with a distinct identity (e.g., an item in stock) and value objects for attributes that describe entities.

### 2. **Microservices Architecture**
   - **Decoupling Services**: Break down your application into smaller, independent services that focus on specific business capabilities (e.g., inventory service, order service).
   - **Scalability**: Microservices can be deployed independently, allowing you to scale parts of the application as needed.
   - **Communication**: Use RESTful APIs, gRPC, or event-driven communication to enable interaction between services.

### 3. **Event-Driven Architecture**
   - **Event Sourcing**: Capture all changes to the state as a sequence of events, allowing you to reconstruct past states or audit history.
   - **CQRS (Command Query Responsibility Segregation)**: Separate the responsibility of handling commands (state changes) from queries (read operations) to optimize performance and scalability.

### 4. **Repository Pattern**
   - **Data Access Layer**: Use repositories to abstract the data access layer, making it easier to swap out data storage solutions or implement caching.
   - **Unit of Work**: Manage transactions efficiently, ensuring all operations on repositories are part of a single transaction.

### 5. **Design Patterns**
   - **Singleton Pattern**: Ensure that certain components (e.g., configuration manager) have a single instance throughout the application's lifecycle.
   - **Factory Pattern**: Use factories to create complex objects, encapsulating the logic required to instantiate them.
   - **Observer Pattern**: Implement observers to allow components to listen and react to events, useful for updating the UI or triggering workflows.

### 6. **Database Design**
   - **Normalization**: Design your database schema to reduce redundancy and improve data integrity.
   - **Indexing**: Optimize queries by indexing columns that are frequently used in search operations.
   - **NoSQL Options**: Consider NoSQL databases if your system requires high scalability or handles unstructured data.

### 7. **Security**
   - **Authentication and Authorization**: Implement secure user authentication (e.g., OAuth2, JWT) and role-based access control.
   - **Data Protection**: Encrypt sensitive data and ensure secure data transmission using protocols like HTTPS.

### 8. **Testing and Continuous Integration**
   - **Automated Testing**: Develop unit, integration, and end-to-end tests to ensure code quality and reliability.
   - **CI/CD Pipelines**: Set up continuous integration and deployment pipelines to automate testing and deployment processes.

### 9. **User Interface/Experience (UI/UX)**
   - **Responsive Design**: Create a user interface that works well on various devices and screen sizes.
   - **Usability Testing**: Conduct usability testing to gather feedback and improve the user experience.

### 10. **Scalability and Performance**
   - **Load Balancing**: Distribute incoming requests evenly across multiple servers to improve response times and availability.
   - **Caching**: Use caching mechanisms (e.g., Redis, Memcached) to reduce load on the database and improve performance for frequently accessed data.

These patterns and practices provide a solid foundation for building a robust inventory management system. Tailor them to your specific needs and constraints, and consider involving stakeholders in the design process to ensure the system aligns with business requirements.

Certainly! Here are some resources and concepts that delve into the specific aspects of managing stock, including registration, consistency, and grouping:

### Key Concepts for Inventory Management Systems

1. **Stock Registration and Tracking**
   - **Unique Identifiers**: Assign a unique identifier (SKU) to each item for easy tracking and retrieval.
   - **Barcode/RFID Integration**: Use barcodes or RFID tags for efficient scanning and updating of stock levels.
   - **Real-time Updates**: Implement a system that updates inventory levels in real-time to reflect storage and retrieval actions accurately.

2. **Consistency and Accuracy**
   - **Atomic Transactions**: Ensure all stock changes (additions, removals) are atomic to prevent data inconsistencies.
   - **Audit Trails**: Maintain logs of all inventory transactions for auditing and accountability.
   - **Periodic Reconciliation**: Conduct regular physical counts to reconcile with system records and adjust discrepancies.

3. **Automatic Retrieval and Restocking**
   - **Threshold Alerts**: Set minimum stock levels to trigger automatic restocking orders when inventory falls below a certain point.
   - **Supplier Integration**: Integrate with supplier systems for seamless order placement and tracking.
   - **Batch Processing**: Implement batch processing for periodic tasks such as inventory checks and order generation.

4. **Grouping and Categorization**
   - **Attributes and Tags**: Use attributes to categorize items (e.g., size, color, quality) and tags to facilitate grouping.
   - **Hierarchical Grouping**: Implement a hierarchy for categories (e.g., electronics > computers > laptops) to organize products logically.
   - **Flexible Queries**: Allow users to query inventory based on multiple attributes to find items meeting specific criteria.

5. **Demand Forecasting and Inventory Optimization**
   - **Historical Data Analysis**: Use historical sales data to predict future demand and adjust stock levels accordingly.
   - **Safety Stock Calculation**: Maintain safety stock to buffer against unforeseen demand spikes or supply delays.

### Recommended Resources

1. **Books**
   - **"Inventory Management Explained: A focus on Forecasting, Lot Sizing, Safety Stock, and Ordering Systems" by David J. Piasecki**: Offers detailed insights into maintaining optimal inventory levels and implementing effective ordering systems.
   - **"Warehouse & Distribution Science" by John Bartholdi and Steven Hackman**: Provides a comprehensive look at warehouse operations, including stock management.

2. **Articles and Papers**
   - **"Inventory Management: Principles, Concepts and Techniques"**: This article covers key inventory management concepts and techniques, including stock registration, consistency, and retrieval.
   - **"The Essentials of Inventory Management" by Max Muller**: Focuses on methods and strategies to keep inventory levels consistent and efficient.

3. **Online Resources**
   - **Supply Chain Management Review**: Offers articles and case studies on inventory management strategies and technologies.
   - **APICS (Association for Supply Chain Management)**: Provides articles, webinars, and resources focused on best practices in inventory management and supply chain optimization.

4. **Software Tools**
   - **Odoo Inventory**: An open-source inventory management solution with features for stock tracking, automated workflows, and reporting.
   - **Zoho Inventory**: Offers tools for managing inventory, order processing, and tracking shipments with robust reporting features.

These resources and concepts should provide a more targeted perspective on building a stock management system that meets your requirements for consistency, automatic retrieval, and flexible grouping.

Understood! To create a command-line application that automatically updates the stock of pills based on your daily intake without requiring manual updates every day, we can implement the following features:

### Key Features
1. **Daily Stock Update**: Automatically update the stock based on your daily intake.
2. **Prescription Tracking**: Input your prescriptions for the week or month, and the software will manage daily updates.
3. **Alerts**: Notify you when stock is low based on your remaining pills and daily requirements.
4. **Persistent Storage**: Use a file (like JSON) to store your medications and prescriptions so that the data persists between runs.
5. **Daily Summary**: Provide a summary of today's requirements and available stock when you run the application.

### Updated Design

#### Data Model
We will keep the same classes but ensure that the `Medication` class has methods to handle daily updates automatically.

```python
class ActivePrinciple:
    def __init__(self, name, daily_dosage):
        self.name = name                 # e.g., "Ibuprofen"
        self.daily_dosage = daily_dosage # e.g., "300mg/day"

class Medication:
    def __init__(self, name, box_size, active_principle, mg_per_pill):
        self.name = name                    # e.g., "Pain Relief"
        self.box_size = box_size            # Number of pills in a box
        self.active_principle = active_principle  # ActivePrinciple object
        self.mg_per_pill = mg_per_pill      # e.g., 150mg
        self.stock = 0                      # Total pills in stock

    def purchase(self, boxes):
        self.stock += boxes * self.box_size

    def dispense(self, amount):
        if self.stock >= amount:
            self.stock -= amount
        else:
            print(f"Not enough stock for {self.name}. Current stock: {self.stock}")

    def calculate_active_units(self):
        return self.stock * self.mg_per_pill

    def is_low_stock(self, threshold):
        return self.stock < threshold

class Prescription:
    def __init__(self, active_principle, days):
        self.active_principle = active_principle  # ActivePrinciple object
        self.days = days                          # Number of days for the prescription
```

### Command-Line Interface

The CLI will allow you to:
1. Add medications and active principles.
2. Add prescriptions for multiple days.
3. Update stock based on daily intake.
4. Check current stock and get alerts.

### Implementation

Here's a complete example of how the main application could look:

```python
import argparse
import json
from datetime import datetime, timedelta

class ActivePrinciple:
    def __init__(self, name, daily_dosage):
        self.name = name                 # e.g., "Ibuprofen"
        self.daily_dosage = daily_dosage # e.g., "300mg/day"

class Medication:
    def __init__(self, name, box_size, active_principle, mg_per_pill):
        self.name = name                    # e.g., "Pain Relief"
        self.box_size = box_size            # Number of pills in a box
        self.active_principle = active_principle  # ActivePrinciple object
        self.mg_per_pill = mg_per_pill      # e.g., 150mg
        self.stock = 0                      # Total pills in stock

    def purchase(self, boxes):
        self.stock += boxes * self.box_size

    def dispense(self, amount):
        if self.stock >= amount:
            self.stock -= amount
        else:
            print(f"Not enough stock for {self.name}. Current stock: {self.stock}")

    def calculate_active_units(self):
        return self.stock * self.mg_per_pill

    def is_low_stock(self, threshold):
        return self.stock < threshold

class Prescription:
    def __init__(self, active_principle, days):
        self.active_principle = active_principle  # ActivePrinciple object
        self.days = days                          # Number of days for the prescription

def load_data():
    """Load medication data from a JSON file."""
    try:
        with open('medications.json', 'r') as f:
            data = json.load(f)
            medications = [Medication(**med) for med in data['medications']]
            prescriptions = [Prescription(ActivePrinciple(**pres['active_principle']), pres['days']) for pres in data['prescriptions']]
            return medications, prescriptions
    except FileNotFoundError:
        return [], []

def save_data(medications, prescriptions):
    """Save medication and prescription data to a JSON file."""
    data = {
        'medications': [med.__dict__ for med in medications],
        'prescriptions': [{
            'active_principle': pres
```
I apologize for not fully detailing the daily stock update functionality in the previous messages. Let's clarify how to implement a **Daily Stock Update** feature in your command-line application. This feature will automatically deduct the required amount of medication from the stock based on your daily prescriptions.

### Detailed Implementation for Daily Stock Update

To implement the daily stock update, we will do the following:

1. **Store the Daily Prescription**: Allow the user to input how many days their prescription lasts and the daily dosage.
2. **Calculate Daily Intake**: Every time you run the application, it will check the current date and determine how many pills you should have taken since the last update.
3. **Update the Stock**: Automatically deduct the required number of pills based on your daily dosage.
4. **Generate Alerts**: Notify you if the stock is low after updating.

### Complete Implementation

Here’s how the application could be structured to implement the daily stock update feature:

```python
import argparse
import json
from datetime import datetime, timedelta

class ActivePrinciple:
    def __init__(self, name, daily_dosage):
        self.name = name                 # e.g., "Ibuprofen"
        self.daily_dosage = daily_dosage # e.g., "300mg/day"

class Medication:
    def __init__(self, name, box_size, active_principle, mg_per_pill):
        self.name = name                    # e.g., "Pain Relief"
        self.box_size = box_size            # Number of pills in a box
        self.active_principle = active_principle  # ActivePrinciple object
        self.mg_per_pill = mg_per_pill      # e.g., 150mg
        self.stock = 0                      # Total pills in stock

    def purchase(self, boxes):
        self.stock += boxes * self.box_size

    def dispense(self, amount):
        if self.stock >= amount:
            self.stock -= amount
        else:
            print(f"Not enough stock for {self.name}. Current stock: {self.stock}")

    def calculate_active_units(self):
        return self.stock * self.mg_per_pill

    def is_low_stock(self, threshold):
        return self.stock < threshold

class Prescription:
    def __init__(self, active_principle, days):
        self.active_principle = active_principle  # ActivePrinciple object
        self.days = days                          # Number of days for the prescription
        self.start_date = datetime.now().date()   # Default to today

def load_data():
    """Load medication data from a JSON file."""
    try:
        with open('medications.json', 'r') as f:
            data = json.load(f)
            medications = [Medication(**med) for med in data['medications']]
            prescriptions = [Prescription(ActivePrinciple(**pres['active_principle']), pres['days']) for pres in data['prescriptions']]
            for pres in prescriptions:
                pres.start_date = datetime.strptime(pres.start_date, '%Y-%m-%d').date()  # Convert to date
            return medications, prescriptions
    except FileNotFoundError:
        return [], []

def save_data(medications, prescriptions):
    """Save medication and prescription data to a JSON file."""
    data = {
        'medications': [med.__dict__ for med in medications],
        'prescriptions': [{
            'active_principle': pres.active_principle.__dict__,
            'days': pres.days,
            'start_date': pres.start_date.strftime('%Y-%m-%d')  # Store as string
        } for pres in prescriptions]
    }
    with open('medications.json', 'w') as f:
        json.dump(data, f)

def update_daily_stock(medications, prescriptions):
    """Update the medication stock based on daily intake."""
    for pres in prescriptions:
        active_principle = pres.active_principle.name
        daily_dosage = int(pres.active_principle.daily_dosage.split()[0])  # e.g., "300mg/day" -> 300
        days_passed = (datetime.now().date() - pres.start_date).days

        # Calculate total pills needed
        total_pills_needed = days_passed * (daily_dosage // medications[0].mg_per_pill)

        for med in medications:
            if med.active_principle.name == active_principle:
                med.dispense(total_pills_needed)
                print(f"Updated stock for {med.name}: {med.stock} pills remaining after taking {total_pills_needed} pills.")

def check_low_stock(medications):
    """Check for low stock and alert the user."""
    threshold = 5  # example threshold for low stock
    for med in medications:
        if med.is_low_stock(threshold):
            print(f"Alert: Stock for {med.name} is low ({
```
To accommodate the scenario where multiple medications can contain the same active principle, we need to ensure that the daily stock update functionality correctly aggregates the dosages from all medications that share the same active principle.

### Updated Implementation

Here’s how we can achieve this in the command-line application:

1. **Aggregate Daily Dosage**: When calculating daily requirements, sum up the dosages from all medications that contain the same active principle.
2. **Update Stock**: Deduct the total required pills across all medications that share an active principle.

### Complete Code Implementation

Below is the updated code that handles multiple medications with the same active principle, including daily stock updates:

```python
import argparse
import json
from datetime import datetime, timedelta

class ActivePrinciple:
    def __init__(self, name, daily_dosage):
        self.name = name                 # e.g., "Ibuprofen"
        self.daily_dosage = daily_dosage # e.g., "300mg/day"

class Medication:
    def __init__(self, name, box_size, active_principle, mg_per_pill):
        self.name = name                    # e.g., "Pain Relief"
        self.box_size = box_size            # Number of pills in a box
        self.active_principle = active_principle  # ActivePrinciple object
        self.mg_per_pill = mg_per_pill      # e.g., 150mg
        self.stock = 0                      # Total pills in stock

    def purchase(self, boxes):
        self.stock += boxes * self.box_size

    def dispense(self, amount):
        if self.stock >= amount:
            self.stock -= amount
        else:
            print(f"Not enough stock for {self.name}. Current stock: {self.stock}")

    def calculate_active_units(self):
        return self.stock * self.mg_per_pill

    def is_low_stock(self, threshold):
        return self.stock < threshold

class Prescription:
    def __init__(self, active_principle, days):
        self.active_principle = active_principle  # ActivePrinciple object
        self.days = days                          # Number of days for the prescription
        self.start_date = datetime.now().date()   # Default to today

def load_data():
    """Load medication data from a JSON file."""
    try:
        with open('medications.json', 'r') as f:
            data = json.load(f)
            medications = [Medication(**med) for med in data['medications']]
            prescriptions = [
                Prescription(ActivePrinciple(**pres['active_principle']), pres['days']) for pres in data['prescriptions']
            ]
            for pres in prescriptions:
                pres.start_date = datetime.strptime(pres.start_date, '%Y-%m-%d').date()  # Convert to date
            return medications, prescriptions
    except FileNotFoundError:
        return [], []

def save_data(medications, prescriptions):
    """Save medication and prescription data to a JSON file."""
    data = {
        'medications': [med.__dict__ for med in medications],
        'prescriptions': [{
            'active_principle': pres.active_principle.__dict__,
            'days': pres.days,
            'start_date': pres.start_date.strftime('%Y-%m-%d')  # Store as string
        } for pres in prescriptions]
    }
    with open('medications.json', 'w') as f:
        json.dump(data, f)

def update_daily_stock(medications, prescriptions):
    """Update the medication stock based on daily intake."""
    for pres in prescriptions:
        active_principle_name = pres.active_principle.name
        daily_dosage = int(pres.active_principle.daily_dosage.split()[0])  # e.g., "300mg/day" -> 300
        days_passed = (datetime.now().date() - pres.start_date).days
        
        # Calculate total pills needed
        total_units_needed = daily_dosage * days_passed  # Total active principle units needed
        total_pills_needed = total_units_needed // next(mg_per_pill for med in medications if med.active_principle.name == active_principle_name)

        # Dispense pills from all medications containing the active principle
        total_dispensed = 0
        for med in medications:
            if med.active_principle.name == active_principle_name:
                if total_dispensed < total_pills_needed:
                    pills_to_dispense = min(total_pills_needed - total_dispensed, med.stock)
                    med.dispense(pills_to_dispense)
                    total_dispensed += pills_to_dispense
                    print(f"Updated stock for {med.name}: {med.stock} pills remaining after taking {pills_to_dispense} pills.")

        if total_dispensed < total_pills_needed:
            print(f"Warning: Not enough total stock for {active
```
Certainly! To enhance the safety and usability of your medication tracking application, we can implement a feature that allows you to record daily pill consumption directly. This way, you won't need to calculate the total dispensed every time you run the application. Instead, you'll simply log how many pills you took each day, and the application will update the stock accordingly.

### Updated Design for Daily Consumption Logging

Here’s how we can modify the application:

1. **Daily Consumption Logging**: You will record your daily consumption of pills. This way, you can track how many pills you've taken without recalculating from the start of the prescription.
2. **Persistent Storage**: Store the consumption history in the JSON file, so you have a record of what you took on each day.
3. **Daily Summary**: Provide a summary of today's requirements and the remaining stock when you run the application.

### Implementation

Below is the updated code that includes daily consumption logging:

```python
import json
from datetime import datetime

class ActivePrinciple:
    def __init__(self, name, daily_dosage):
        self.name = name                 # e.g., "Ibuprofen"
        self.daily_dosage = daily_dosage # e.g., "300mg/day"

class Medication:
    def __init__(self, name, box_size, active_principle, mg_per_pill):
        self.name = name                            # e.g., "Pain Relief"
        self.box_size = box_size                    # Number of pills in a box
        self.active_principle = active_principle    # ActivePrinciple object
        self.mg_per_pill = mg_per_pill              # e.g., 150mg
        self.stock = 0                              # Total pills in stock

    def purchase(self, boxes):
        self.stock += boxes * self.box_size

    def dispense(self, amount):
        if self.stock >= amount:
            self.stock -= amount
        else:
            print(f"Not enough stock for {self.name}. Current stock: {self.stock}")

    def calculate_active_units(self):
        return self.stock * self.mg_per_pill

    def is_low_stock(self, threshold):
        return self.stock < threshold

class Prescription:
    def __init__(self, active_principle, days):
        self.active_principle = active_principle  # ActivePrinciple object
        self.days = days                          # Number of days for the prescription
        self.start_date = datetime.now().date()   # Default to today
        self.consumption_log = {}                 # Dictionary to track daily consumption

def load_data():
    """Load medication data from a JSON file."""
    try:
        with open('medications.json', 'r') as f:
            data = json.load(f)
            medications = [Medication(**med) for med in data['medications']]
            prescriptions = [
                Prescription(ActivePrinciple(**pres['active_principle']), pres['days']) for pres in data['prescriptions']
            ]

            # Load consumption log
            for pres in prescriptions:
                pres.start_date = datetime.strptime(pres.start_date, '%Y-%m-%d').date()  # Convert to date
                pres.consumption_log = {datetime.strptime(date, '%Y-%m-%d').date(): count for date, count in pres.consumption_log.items()}
                pres.consumption_log = pres.consumption_log or {}  # Initialize if empty

            return medications, prescriptions
    except FileNotFoundError:
        return [], []

def save_data(medications, prescriptions):
    """Save medication and prescription data to a JSON file."""
    data = {
        'medications': [med.__dict__ for med in medications],
        'prescriptions': [{
            'active_principle': pres.active_principle.__dict__,
            'days': pres.days,
            'start_date': pres.start_date.strftime('%Y-%m-%d'),  # Store as string
            'consumption_log': pres.consumption_log
        } for pres in prescriptions]
    }
    with open('medications.json', 'w') as f:
        json.dump(data, f)

def log_consumption(prescriptions, med_name, amount):
    """Log the daily consumption of pills."""
    for pres in prescriptions:
        for med in medications:
            if med.name == med_name and med.active_principle.name == pres.active_principle.name:
                if datetime.now().date() not in pres.consumption_log:
                    pres.consumption_log[datetime.now().date()] = 0
                pres.consumption_log[datetime.now().date()] += amount
                med.dispense(amount)  # Update stock based on consumption
                print(f"Logged {amount} pills for {med.name}. Remaining stock: {med.stock}")
                return
    print(f"Medication {med_name} not found in prescriptions.")

def daily_summary(medications, prescriptions):
    """Print a summary of today's stock and consumption."""
    print("\nDaily Summary:")
    for pres in prescriptions:
        active_principle_name = pres.active_principle.name
        total_consumed = pres.consumption_log.get(datetime.now().date(), 0)
        print(f"Active Principle: {active_principle_name}, Total Consumed Today: {total_consumed} mg")

        # Check stock for medications containing this active principle
        for med in medications:
            if med.active_principle.name == active_principle_name:
                print(f"- {med.name}: {med.stock} pills remaining ({med.calculate_active_units()} mg available)")

def main():
    medications, prescriptions = load_data()

    # Example: Log consumption of 2 pills for "Pain Relief"
    log_consumption(prescriptions, "Pain Relief", 2)

    # Example: Print daily summary
    daily_summary(medications, prescriptions)

    save_data(medications, prescriptions)

if __name__ == "__main__":
    main()
```

Certainly! Below is a workflow diagram for the medication tracking application. The diagram outlines the main processes involved in managing medications, prescriptions, and daily consumption logging. 

### Workflow Diagram

```plaintext
+-----------------------------------+
|          Start Application        |
+-----------------+-----------------+
                  |
                  v
          +-------+-------+
          | Load Data     |
          | from JSON     |
          +-------+-------+
                  |
                  v
          +-------+-------+
          | Display Menu  |
          +-------+-------+
                  |
                  v
          +-------+-------+
          | User Input    |A
          | (Choose Action)|
          +-------+-------+
                  |
                  v
      +-----------+-----------+
      |           |           |
      v           v           v
+-----+-----+ +---+---+ +-----+-----+
| Add Active| | Add Med| | Log Daily |
| Principle  | | ication| | Consumption|
+-----+-----+ +---+---+ +-----+-----+
      |           |           |
      |           |           |
      |           |           v
      |           |    +------+------+
      |           |    | Record Daily |
      |           |    | Consumption  |
      |           |    +------+------+
      |           |           |
      |           v           |
      |        +--------------+--------------+
      |        | Update Stock Based on Log   |
      |        +--------------+--------------+
      |                       |
      |                       v
      |                +------+------+
      |                | Check Low   |
      |                | Stock       |
      |                +------+------+
      |                       |
      |                       v
      |                +------+------+
      |                | Display Alert |
      |                +------+------+
      |                       |
      |                       v
      |                +------+------+
      |                | Daily Summary|
      |                +------+------+
      |                       |
      |                       v
      |                +------+------+
      |                | Save Data to |
      |                | JSON File    |
      |                +------+------+
      |                       |
      |                       v
      +-------------------> +------+------+
      |                     | Exit          |
      +-------------------> +--------------+
```

### Description of Workflow Steps

1. **Start Application**: The user starts the application.
2. **Load Data from JSON**: The application attempts to load existing medication and prescription data from a JSON file.
3. **Display Menu**: The main menu is displayed, allowing the user to choose an action.
4. **User Input (Choose Action)**: The user selects an action such as adding an active principle, adding a medication, logging daily consumption, or viewing a summary.
5. **Add Active Principle**: If the user opts to add an active principle, the application collects the necessary data and updates the internal data structure.
6. **Add Medication**: If the user opts to add a medication, the application collects the necessary data and updates the internal data structure.
7. **Log Daily Consumption**: If the user opts to log daily consumption, the application collects the medication name and quantity consumed.
8. **Record Daily Consumption**: The application updates the consumption log for the corresponding prescription and medication.
9. **Update Stock Based on Log**: The application deducts the consumed amount from the current stock of the medication.
10. **Check Low Stock**: The application checks if the stock of any medication is below a predefined threshold.
11. **Display Alert**: If any medication is low on stock, the application alerts the user.
12. **Daily Summary**: The application provides a summary of today's consumption and current stock levels for each medication.
13. **Save Data to JSON File**: The updated data for medications and prescriptions, including consumption logs, is saved back to the JSON file.
14. **Exit**: The application exits.

### Conclusion

This diagram and the associated descriptions provide a comprehensive overview of how the medication tracking application operates. It outlines the key processes and interactions involved in managing your medications and daily intake effectively. If you need any further adjustments or additional details, feel free to ask!

To allow for updating existing prescriptions and adding new prescriptions without affecting past data, we can enhance the medication tracking application with the following features:

1. **Update Existing Prescription**: Allow the user to update a prescription for an active principle but specify that changes take effect from a certain date.
2. **Add New Prescriptions**: Allow for the addition of new prescriptions that can coexist with existing ones.

### Implementation Steps

1. **Store Prescription Start Dates**: Each prescription will have a start date and an end date, and any updates will apply from that date onward.
2. **Modify Log Consumption**: Ensure that the log for daily consumption only affects prescriptions that are active (i.e., those whose start date is on or before the current date).
3. **User Commands**: Implement command-line options to update existing prescriptions and add new ones.