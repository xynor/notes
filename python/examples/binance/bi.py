from binance.client import Client
from binance.exceptions import BinanceAPIException, BinanceWithdrawException

api_key = 'x38FyretN1TTGkWnvlEn1DUOSTLjjoftiNaBKXpt9TI25eCjheNr1X1MkBvPBSmR'
api_secret = 'PoqkAM8pxCQvMzPRQ9WkwB22uztWfz11bvudyjkM6C0Vq2UabfvFdhwnpBkXWCIQ'
client = Client(api_key, api_secret)
# his = client.get_withdraw_history()
# print(his)
# depth = client.get_order_book(symbol='BNBBTC')
# print(depth)
# withdraw 100 ETH
# check docs for assumptions around withdrawals
try:
	result = client.withdraw(
		asset='USDT',
		address='TD42q2KbXtjokFPAvhzzYXpucPqSm6vEqm',
		amount=2,
		recvWindow=5000,
		network='TRC20',
	)
except BinanceAPIException as e:
	print(e)
except BinanceWithdrawException as e:
	print(e)
else:
	print("Success:", result)
