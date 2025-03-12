Feature: Seller appoints location for order

  Background:
    Given an order with ID "67d150af6b6922ff40714555" is created
    And an appointment with ID "5a3f3e5a8b15c19c9aebed4e" is created for this order

  Scenario: Seller appoints a location for order and buyer rejects it
    When the seller chooses a location for order with ID "67d150af6b6922ff40714555"
    And the buyer rejects the seller's location for order with ID "67d150af6b6922ff40714555"
    Then the order status for order with ID "67d150af6b6922ff40714555" should be 0
    And the response status should be 200

  Scenario: Seller appoints a location for order and buyer accepts it
    When the seller chooses a location for order with ID "67d150af6b6922ff40714555"
    And the buyer accepts the seller's location for order with ID "67d150af6b6922ff40714555"
    Then the order status for order with ID "67d150af6b6922ff40714555" should be 2
    And the response status should be 200

  Scenario: Seller changes location for an order
    When the seller updates the location for order with ID "67d150af6b6922ff40714555"
    Then the order status for order with ID "67d150af6b6922ff40714555" should not change
    And the response status should be 200
