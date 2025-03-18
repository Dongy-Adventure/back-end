Feature: Buyer confirms an order

  As a buyer  
  I want to confirm my order  
  So that I can proceed to payment and finalize my purchase  

  Background:  
    Given a buyer with ID "67d150af6b6922ff40714555"
    And a product with ID "67d151ba6b6922ff40714558" exists in the buyer cart  

  Scenario: Successful order confirmation and payment  
    When the payment status is "success"
    Then an order should be created with respone status 201  
    And the product should be removed from the buyer cart

  Scenario: Failed payment process  
    When the payment status is "failure"  
    And a log should be generated indicating "Payment failure: order not created"

  Scenario: Buyer cancels the order before payment  
    When the payment status is "canceled"
    And a log should be generated indicating "Order cancelled by the buyer"

  Scenario: Order confirmation with insufficient stock  
    And the product is out of stock  
    And a log should be generated indicating "Insufficient stock for product"
